package user

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest
// Query binding 需要打 form 标签，Body json binding 需要打 json 标签
type RegisterRequest struct {
	Username       string  `json:"username" binding:"required"`
	Password       string  `json:"password" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	InvitationCode *string `json:"invitation_code" binding:"omitempty"`
}

type RegisterDataResponse struct {
	Email    string `json:"email"`
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

// Register 注册 (POST /register)
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	// 检查是否允许注册
	if !config.RegisterConfig.AllowRegister {
		resp.Abort(c, code.UserErrorRegisterNotAllowed)
		return
	}

	err := util.CheckUsername(req.Username)
	if err != nil {
		resp.AbortWithMsg(c, code.UserErrorInvalidUsername, err.Error())
		return
	}

	var inviterID int32 = 0

	// 无邀请码注册，检查是否允许无邀请码注册
	if req.InvitationCode == nil || *req.InvitationCode == "" {
		if config.RegisterConfig.UseInvitationCode {
			resp.AbortWithMsg(c, code.UserErrorInvalidInvitationCode, "invitation code is required")
			return
		}
	} else {
		// 邀请码功能, 有邀请码注册，检查邀请码是否有效
		inviterID, err = cache.GetInviterIDByInvitationCode(*req.InvitationCode)
		if err != nil {
			resp.AbortWithMsg(c, code.UserErrorInvalidInvitationCode, "invalid invitation code")
			log.Logger.Error("invalid invitation code: " + err.Error())
			return
		}
		log.Logger.Info("invitation code: " + *req.InvitationCode)
	}

	// 检查用户是否已存在
	if db.CheckUserExist(req.Username, req.Email) {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordCreateFailed, "user already exists")
		log.Logger.Error("failed to register: user already exists")
		return
	}

	// 生成密码哈希
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, "failed to hash password")
		log.Logger.Error("failed to hash password: " + err.Error())
		return
	}

	// 注册
	err = db.CreateUser(&model.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   string(password),
		LastActive: time.Now(),
		Inviter:    inviterID,
	})
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordCreateFailed, "failed to register "+err.Error())
		log.Logger.Error("failed to register: " + err.Error())
		return
	}

	// 获取用户注册的 userID
	user, err := db.GetUserByEmail(req.Email)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "failed to get user by email")
		log.Logger.Error("failed to get user by email: " + err.Error())
		return
	}

	// 消费邀请码
	if req.InvitationCode != nil && *req.InvitationCode != "" {
		err = cache.ConsumeInvitationCode(*req.InvitationCode, user.UserID)
		if err != nil {
			resp.AbortWithMsg(c, code.UnknownError, "failed to consume invitation code")
			return
		}
	}

	resp.OKWithData(c, &RegisterDataResponse{
		Email:    user.Email,
		UserID:   user.UserID,
		Username: user.Username,
	})
	log.Logger.Infof("register success, userID: %v", user.UserID)
}

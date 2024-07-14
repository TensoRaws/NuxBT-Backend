package user

import (
	"strconv"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/service/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
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
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

// Register 注册 (POST /register)
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.AbortWithMsg(c, "invalid request")
		return
	}

	// 无邀请码注册，检查是否允许无邀请码注册
	if req.InvitationCode == nil || *req.InvitationCode == "" {
		if config.ServerConfig.UseInvitationCode {
			util.AbortWithMsg(c, "invitation code is required")
			return
		}
	} else {
		// 有邀请码注册，检查邀请码是否有效
		// do something
		// 未实现
		// OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
		log.Logger.Info("invitation code: ", *req.InvitationCode)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.AbortWithMsg(c, "failed to hash password")
		log.Logger.Error("failed to hash password: " + err.Error())
		return
	}
	// 注册
	err = dao.CreateUser(&model.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   string(password),
		LastActive: time.Now(),
	})
	if err != nil {
		util.AbortWithMsg(c, "failed to register: ")
		log.Logger.Error("failed to register: " + err.Error())
		return
	}

	user, err := dao.GetUserByEmail(req.Email)
	if err != nil {
		util.AbortWithMsg(c, "failed to get user by email")
		log.Logger.Error("failed to get user by email: " + err.Error())
		return
	}

	util.OKWithData(c, RegisterDataResponse{
		Email:    user.Email,
		UserID:   strconv.FormatInt(int64(user.UserID), 10),
		Username: user.Username,
	})
	log.Logger.Info("register success: " + util.StructToString(user))
}

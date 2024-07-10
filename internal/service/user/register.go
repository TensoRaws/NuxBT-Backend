package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

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

func Register(c *gin.Context) {
	log.Logger.Debug(c.Query("username"), c.Query("password"), c.Query("email"), c.Query("invitation_code"))

	var req RegisterRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger.Debug("invalid request: ", err.Error())
		util.AbortWithMsg(c, "invalid request")
		return
	}

	// 无邀请码注册，检查是否允许无邀请码注册
	if req.InvitationCode == nil {
		if config.GetString("register.useInvitationCode") == "true" {
			util.AbortWithMsg(c, "invitation code is required")
		}
	} else {
		// 有邀请码注册，检查邀请码是否有效
		// do something
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.AbortWithMsg(c, "failed to hash password")
		log.Logger.Error("failed to hash password: ", err.Error())
		return
	}
	// 注册
	err = CreateUser(&model.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   string(password),
		LastActive: time.Now(),
	})
	if err != nil {
		util.AbortWithMsg(c, "failed to register: ")
		log.Logger.Error("failed to register: ", err.Error())
		return
	}

	user, err := GetUserByEmail(req.Email)
	if err != nil {
		util.AbortWithMsg(c, "failed to get user by email")
		log.Logger.Error("failed to get user by email: ", err.Error())
		return
	}
	util.OKWithDataStruct(c, RegisterDataResponse{
		Email:    user.Email,
		UserID:   string(user.UserID),
		Username: user.Username,
	})

	return
}

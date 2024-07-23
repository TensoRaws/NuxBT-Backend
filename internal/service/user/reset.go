package user

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

// ResetPassword 重置密码 (POST /password/reset)
func ResetPassword(c *gin.Context) {
	// 绑定参数
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, "failed to hash password")
		log.Logger.Error("failed to hash password: " + err.Error())
		return
	}

	// 修改密码
	err = db.PatchUser(userID, &model.User{Password: string(password)})

	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordPatchFailed, "reset password fail")
	}
	// 返回
	resp.OK(c)
	log.Logger.Infof("Reset password success, user ID: %v", userID)
}

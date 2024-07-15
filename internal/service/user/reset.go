package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
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

	user, err := dao.GetUserByID(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "User not found")
		return
	}

	// 修改密码
	err = dao.UpdateUserPassword(user, req.NewPassword)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordUpdateFailed, "reset password fail")
	}
	// 返回
	resp.OK(c)
	log.Logger.Info("Reset password success: " + util.StructToString(user))
}

package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
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
		util.AbortWithMsg(c, "invalid request")
		return
	}

	// 鉴权
	userID, err := util.GetUserIDFromGinContext(c)
	if err != nil {
		util.AbortWithMsg(c, "Please login first")
		return
	}

	user, err := dao.GetUserByID(userID)
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	// 修改密码
	err = dao.SetUserPassword(user, req.NewPassword)
	if err != nil {
		util.AbortWithMsg(c, "reset password fail")
	}
	// 返回
	util.OKWithMsg(c, "reset password success")
	log.Logger.Info("Reset password success: " + util.StructToString(user))
}

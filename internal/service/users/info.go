package users

import (
	"github.com/TensoRaws/NuxBT-Backend/module/response"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// Info /douyin/users/ - 用户信息
func Info(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	userinfo, err := response.GetUserResponseByOwnerId(userID)
	if err != nil {
		response.ErrRespWithMsg(c, "User not found")
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"users": *userinfo,
	})
}

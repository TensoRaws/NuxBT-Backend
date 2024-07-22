package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type RoleResponse struct {
	Roles []string `json:"roles"`
}

// RoleMe 获取用户自己的角色信息 (GET /role/me)
func RoleMe(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	roles, err := cache.GetUserRolesByID(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Info("Failed to get user roles: " + err.Error())
		return
	}

	resp.OKWithData(c, &RoleResponse{
		Roles: roles,
	})

	log.Logger.Infof("get user role success, userID: %v", userID)
}

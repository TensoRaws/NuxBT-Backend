package user

import (
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// Logout 用户登出 (POST /logout)
func Logout(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	resp.OK(c)

	log.Logger.Info("Logout success, user ID: " + util.StructToString(userID))
}

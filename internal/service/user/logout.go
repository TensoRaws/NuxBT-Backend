package user

import (
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// Logout 用户登出 (POST /logout)
func Logout(c *gin.Context) {
	user, err := util.GetUserIDFromGinContext(c)
	if err != nil {
		util.AbortWithMsg(c, "Please login first")
		return
	}

	util.OKWithMsg(c, "Logout success")

	log.Logger.Info("Logout success: " + util.StructToString(user))
}

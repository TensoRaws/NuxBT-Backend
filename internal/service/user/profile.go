package user

import (
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// ProfileMe 获取用户自己的信息 (GET /profile/me)
func ProfileMe(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	util.OKWithMsg(c, "User found"+strconv.FormatInt(userID, 10))
}

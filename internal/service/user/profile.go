package user

import (
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ProfileMe 获取用户自己的信息 (GET /user/profile/me)
func ProfileMe(c *gin.Context) {
	userID := util.GetUserIDFromGinContext(c)

	util.OKWithMsg(c, "User found"+strconv.FormatInt(userID, 10))
}

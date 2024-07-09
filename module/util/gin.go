package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserIDFromGinContext 从 RequireAuth 处读取 user_id
func GetUserIDFromGinContext(c *gin.Context) int64 {
	userIDstr := c.GetString("user_id")
	// 未登录
	if len(userIDstr) == 0 {
		return 0
	}
	// 已登录
	userID, _ := strconv.ParseInt(userIDstr, 10, 64)

	return userID
}

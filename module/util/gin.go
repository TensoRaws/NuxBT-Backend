package util

import (
	"net/http"
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

func OKWithMsg(c *gin.Context, ok string) {
	resp := map[string]interface{}{
		"success": true,
		"message": ok,
	}
	c.JSON(http.StatusOK, resp)
}

func OKWithData(c *gin.Context, data map[string]interface{}) {
	resp := map[string]interface{}{
		"success": true,
		"message": "ok",
		"data":    data,
	}
	c.JSON(http.StatusOK, resp)
}

func AbortWithMsg(c *gin.Context, msg string) {
	resp := map[string]interface{}{
		"success": false,
		"message": msg,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

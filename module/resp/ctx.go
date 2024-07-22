package resp

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserIDFromGinContext 从 RequireAuth 处读取 user_id
func GetUserIDFromGinContext(c *gin.Context) (int32, error) {
	userIDstr := c.GetString("user_id")
	// 未登录
	if len(userIDstr) == 0 {
		return -1, fmt.Errorf("user_id is null")
	}
	// 已登录
	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	return int32(userID), err
}

// GetRolesFromGinContext 从 RequireAuth 处读取 roles
func GetRolesFromGinContext(c *gin.Context) ([]string, error) {
	roles := c.GetStringSlice("roles")
	if len(roles) == 0 {
		return nil, fmt.Errorf("roles is null")
	}

	return roles, nil
}

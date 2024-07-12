package util

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/gin-gonic/gin"
)

// GetUserIDFromGinContext 从 RequireAuth 处读取 user_id
func GetUserIDFromGinContext(c *gin.Context) (int64, error) {
	userIDstr := c.GetString("user_id")
	// 未登录
	if len(userIDstr) == 0 {
		return -1, fmt.Errorf("user_id is null")
	}
	// 已登录
	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	return userID, err
}

// OKWithMsg 返回成功信息
func OKWithMsg(c *gin.Context, ok string) {
	resp := map[string]interface{}{
		"success": true,
		"message": ok,
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

// OKWithData 返回成功信息，携带自定义数据（结构体）
func OKWithData(c *gin.Context, cache bool, data interface{}) {
	resp := map[string]interface{}{
		"success": true,
		"message": "ok",
		"data":    data,
	}
	if cache {
		c.Set("cache",
			StructToString(
				map[string]interface{}{
					"success": true,
					"message": "cache",
					"data":    data,
				},
			),
		)
	}

	c.JSON(http.StatusOK, resp)
}

// OKWithCache 返回缓存数据，终止请求
func OKWithCache(c *gin.Context, cache string) {
	var resp interface{}
	err := StringToStruct(cache, &resp)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp)
	c.Abort()
}

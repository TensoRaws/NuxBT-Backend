package resp

import (
	"fmt"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"net/http"
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

// OK 返回成功信息
func OK(c *gin.Context) {
	resp := map[string]interface{}{
		"success": true,
	}
	c.JSON(http.StatusOK, resp)
}

// OKWithData 返回成功信息，携带自定义数据（结构体）
func OKWithData(c *gin.Context, data interface{}) {
	resp := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	c.JSON(http.StatusOK, resp)
}

// Abort 返回错误码
func Abort(c *gin.Context, code code.Code) {
	errorResp := map[string]interface{}{
		"code":    code,
		"message": code.String(),
	}
	resp := map[string]interface{}{
		"success": false,
		"error":   errorResp,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

// AbortWithMsg 返回错误码，自定义错误信息
func AbortWithMsg(c *gin.Context, code code.Code, msg string) {
	errorResp := map[string]interface{}{
		"code":    code,
		"message": code.String() + ": " + msg,
	}
	resp := map[string]interface{}{
		"success": false,
		"error":   errorResp,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

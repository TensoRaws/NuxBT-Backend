package jwt

import (
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// RequireAuth 鉴权中间件
// 如果用户携带的 token 验证通过，将 user_id 存入上下文中然后执行下一个 Handler
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token，来自 JWT Blacklist 中间件的 c.Set("token", token)
		token := c.GetString("token")
		claims, err := ParseToken(token)
		if err != nil {
			util.AbortWithMsg(c, "TOKEN IS INVALID, Please Log In")
			return
		}

		userID := claims.ID
		c.Set("user_id", userID)
		// 放行
		c.Next()
	}
}

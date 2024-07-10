package jwt

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// RequireAuth 鉴权中间件
// 如果用户携带的 token 验证通过，将 user_id 存入上下文中然后执行下一个 Handler
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		log.Logger.Infof("url: %v", url.Path)
		// 从输入的 url 中查询 token 值
		token := c.Query("token")
		if len(token) == 0 {
			// 从输入的表单中查询 token 值
			token = c.PostForm("token")
		}

		if len(token) == 0 {
			util.AbortWithMsg(c, "JSON WEB TOKEN IS NULL")
			return
		}

		log.Logger.Info("Get token successfully")
		// auth = [[header][cliams][signature]]
		// 解析 token
		claims, err := ParseToken(token)
		if err != nil {
			util.AbortWithMsg(c, "ERR_INVALID_TOKEN")
			return
		}
		// validate expire time
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			util.AbortWithMsg(c, "TOKEN IS ALREADY EXPIRED, Please Log In Again")
			return
		}

		userId := claims.ID
		c.Set("user_id", userId)
		c.Set("is_login", true)
		// 放行
		c.Next()
	}
}

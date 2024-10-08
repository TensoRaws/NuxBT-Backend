package jwt

import (
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

const JWT_PREFIX = "jwt:"

// RequireAuth 鉴权中间件
// 如果用户携带的 token 验证通过，将 user_id 存入上下文中然后执行下一个 Handler
func RequireAuth(addBlacklist bool) gin.HandlerFunc {
	redisClient := cache.Cache

	return func(c *gin.Context) {
		// 从请求头中获取 token
		token := c.Request.Header.Get("Authorization")

		log.Logger.Info("Get token successfully")

		// 检查 Token 是否存在于 Redis 黑名单中
		exists := redisClient.Exists(JWT_PREFIX + token).Val()
		if exists > 0 {
			log.Logger.Info("Token has been blacklisted")
			resp.Abort(c, code.AuthErrorTokenHasBeenBlacklisted)
			return
		}
		// 如果 Token 不在黑名单中，继续处理请求
		claims, err := ParseToken(token)
		if err != nil {
			resp.AbortWithMsg(c, code.AuthErrorTokenIsInvalid, "Please Log In")
			return
		}
		userID := claims.ID
		// 在上下文中存储 token 和 user_id
		c.Set("token", token)
		c.Set("user_id", userID)
		// 放行
		c.Next()

		// 如果启用拉黑模式，处理请求拉黑 Token
		if addBlacklist {
			err := redisClient.Set(JWT_PREFIX+token, "", GetJWTTokenExpiredDuration()).Err()
			if err != nil {
				log.Logger.Error("Error adding token to blacklist: " + err.Error())
			}
		}
	}
}

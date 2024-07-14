package cache

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTBlacklist 检查JWT是否在黑名单中
func JWTBlacklist(redisClient *cache.Client, addBlacklist bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization
		authHeader := c.Request.Header.Get("Authorization")

		// 检查 Authorization 头部是否存在
		if authHeader == "" {
			util.AbortWithMsg(c, "Authorization header is missing")
			return
		}

		// 检查 Authorization 是否以 "Bearer" 开始
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.AbortWithMsg(c, "Invalid Authorization header format")
			return
		}

		// 第二部分是 Token
		token := parts[1]

		log.Logger.Info("Get token successfully")

		// 检查 Token 是否存在于 Redis 黑名单中
		exists := redisClient.Exists(token).Val()
		if exists > 0 {
			log.Logger.Info("Token has been blacklisted")
			util.AbortWithMsg(c, "Token has been blacklisted")
			return
		}

		// 如果 Token 不在黑名单中，继续处理请求
		c.Set("token", token)
		c.Next()

		// 如果启用拉黑模式，处理请求拉黑 Token
		if addBlacklist {
			err := redisClient.Set(token, "", jwt.GetJWTTokenExpiredDuration()).Err()
			if err != nil {
				log.Logger.Error("Error adding token to blacklist: " + err.Error())
			}
		}
	}
}

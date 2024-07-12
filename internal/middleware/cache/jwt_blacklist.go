package cache

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// JWTBlacklist 检查JWT是否在黑名单中
func JWTBlacklist(redisClient *cache.Client, addBlacklist bool) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		// 检查 Token 是否存在于 Redis 黑名单中
		exists := redisClient.Exists(token).Val()
		if exists > 0 {
			log.Logger.Info("Token has been blacklisted")
			util.AbortWithMsg(c, "Token has been blacklisted")
			return
		}

		// 如果 Token 不在黑名单中，继续处理请求
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

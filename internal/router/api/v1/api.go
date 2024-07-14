package v1

import (
	"net/http"
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/logger"
	user_service "github.com/TensoRaws/NuxBT-Backend/internal/service/user"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件
	r.Use(middleware_cache.NewRateLimiter(cache.Clients[cache.IPLimit],
		config.ServerConfig.RequestLimit, 60*time.Second)) // 限流中间件

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "NuxBT-Backend",
		})
	})

	api := r.Group("/api/v1/")
	{
		user := api.Group("user/")
		{
			// 用户注册
			user.POST("register", user_service.Register)
			// 用户登录
			user.POST("login", user_service.Login)
			// 用户登出
			user.POST("logout",
				jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], true), // 把 token 拉黑
				user_service.Logout,
			)
			// 用户信息
			user.GET("profile/me",
				jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], false),
				user_service.ProfileMe,
			)
			// 修改密码
			user.POST("password/reset",
				jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], true), // 把 token 拉黑
				user_service.ResetPassword)
		}
	}

	return r
}

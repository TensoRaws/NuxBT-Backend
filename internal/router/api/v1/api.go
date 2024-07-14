package v1

import (
	"net/http"
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/cros"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/logger"
	user_service "github.com/TensoRaws/NuxBT-Backend/internal/service/user"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(cros.CorsByRules(config.ServerConfig.Cros)) // 跨域中间件
	r.Use(logger.DefaultLogger(), gin.Recovery())     // 日志中间件
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
			// 用户查询他人信息
			user.GET("profile/",
				jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], false),
				middleware_cache.Response(cache.Clients[cache.RespCache], 1*time.Minute),
				user_service.ProfileOthers,
			)

			// 修改密码
			user.POST("password/reset",
				jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], true), // 把 token 拉黑
				user_service.ResetPassword)
		}
		// 更新用户信息
		user.POST("profile/update",
			jwt.RequireAuth(cache.Clients[cache.JWTBlacklist], false),
			user_service.ChangeUser)
	}

	return r
}

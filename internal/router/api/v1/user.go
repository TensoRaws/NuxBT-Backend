package v1

import (
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/rbac"
	user_service "github.com/TensoRaws/NuxBT-Backend/internal/service/user"
	"github.com/TensoRaws/NuxBT-Backend/module/role"
	"github.com/gin-gonic/gin"
)

// UserRouterGroup 用户路由组
func UserRouterGroup(api *gin.RouterGroup) {
	user := api.Group("user/")

	// 用户注册
	user.POST("register", user_service.Register)
	// 用户登录
	user.POST("login", user_service.Login)
	// 用户刷新 token
	user.POST("token/refresh",
		jwt.RequireAuth(true), // 把 token 拉黑
		user_service.TokenRefresh)
	// 用户登出
	user.POST("logout",
		jwt.RequireAuth(true), // 把 token 拉黑
		user_service.Logout,
	)
	// 修改密码
	user.POST("password/reset",
		jwt.RequireAuth(false),
		user_service.ResetPassword)
	// 用户角色
	user.GET("role/me",
		jwt.RequireAuth(false),
		user_service.RoleMe,
	)
	// 用户信息
	user.GET("profile/me",
		jwt.RequireAuth(false),
		user_service.ProfileMe,
	)
	// 用户查询他人信息
	user.GET("profile/",
		jwt.RequireAuth(false),
		middleware_cache.Response(1*time.Minute),
		user_service.ProfileOthers,
	)
	// 用户信息更新
	user.POST("profile/update",
		jwt.RequireAuth(false),
		user_service.ProfileUpdate)
	// 用户邀请码生成
	user.POST("invitation/gen",
		jwt.RequireAuth(false),
		rbac.RBAC(role.ADVANCED_USER, role.VIP), // 高级用户权限
		user_service.InvitationGen)
	// 用户邀请码列表
	user.GET("invitation/me",
		jwt.RequireAuth(false),
		middleware_cache.Response(5*time.Second),
		user_service.InvitationMe)
}

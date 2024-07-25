package rbac

import (
	cache_logic "github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// RBAC 获取用户角色，存入上下文，进行权限控制，allowRoles 为允许的角色，为空则不限制
func RBAC(allowRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := resp.GetUserIDFromGinContext(c)

		// 从 Redis 获取用户角色
		roles, err := cache_logic.GetUserRolesByID(userID)
		if err != nil {
			// 处理 Redis 查询错误
			log.Logger.Error("RBAC GetUserRolesByID Error: " + err.Error())
			resp.AbortWithMsg(c, code.UnknownError, err.Error())
			return
		}

		if len(allowRoles) > 0 {
			// 检查用户是否有合适的角色
			hasAllowedRole := false
			for _, role := range roles {
				if util.CheckStringInSlice(role, allowRoles) {
					hasAllowedRole = true
					break
				}
			}

			// 用户没有合适的角色，拦截请求
			if !hasAllowedRole {
				resp.AbortWithMsg(c, code.AuthErrorNoPermission, "Role has no permission")
				log.Logger.Errorf("RBAC Role has no permission, userID: %d, roles: %v", userID, roles)
				return
			}
		}

		// 将角色信息存储在 Gin 上下文中
		c.Set("roles", roles)

		// 继续处理请求
		c.Next()
	}
}

package rbac

import (
	cache_logic "github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

// RABC 获取用户角色，存入上下文，进行权限控制
func RABC(allowRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设从JWT或其他方式获取用户ID
		userID, _ := resp.GetUserIDFromGinContext(c)

		// 从Redis获取用户角色
		roles, err := cache_logic.GetUserRolesByID(userID)
		if err != nil {
			// 处理Redis查询错误
			log.Logger.Error("RABC GetUserRolesByID Error: " + err.Error())
			resp.AbortWithMsg(c, code.UnknownError, err.Error())
			return
		}

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
			log.Logger.Errorf("RABC Role has no permission, userID: %d, roles: %v", userID, roles)
			return
		}

		// 将角色信息存储在Gin上下文中
		c.Set("roles", roles)

		// 继续处理请求
		c.Next()
	}
}

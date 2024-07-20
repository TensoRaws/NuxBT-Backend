package cache

import (
	"sort"
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
)

const (
	ROLE_PREFIX = "role:userID:"
)

// GetUserRolesByID 根据 userID 获取用户角色列表
func GetUserRolesByID(userID int32) ([]string, error) {
	c := cache.Cache

	var roles []string
	// 从 Redis 获取用户角色
	roleKey := ROLE_PREFIX + strconv.Itoa(int(userID))

	keyExists, err := c.Exists(roleKey).Result()
	if err != nil {
		log.Logger.Error("Role Unknow Cache Error: " + err.Error())
		return nil, err
	}

	// 如果 Redis 中没有对应的 set，查询数据库
	if keyExists == 0 {
		log.Logger.Info("Role Cache Miss: " + roleKey)
		roles, err = db.GetUserRolesByID(userID)
		if err != nil {
			// 未知错误，抛了，正常情况下不会出现
			log.Logger.Error("Role DB Unknow Error: " + err.Error())
			return nil, err
		}

		// 将数据库中查询到的角色保存到 Redis
		if len(roles) > 0 {
			err = c.SAdd(roleKey, roles).Err()
		} else {
			// 用户没有特殊角色时，插入空字符串，防止 Redis 自动删除 key
			err = c.SAdd(roleKey, "").Err()
		}
		if err != nil {
			log.Logger.Error("Role Unknow Cache Error: " + err.Error())
			return nil, err
		}
		// 设置过期时间
		err = c.Expire(roleKey, cache.DefaultExpiration).Err()
		if err != nil {
			log.Logger.Error("Role Unknow Cache Error: " + err.Error())
			return nil, err
		}
	} else {
		log.Logger.Info("Role Cache Hit: " + roleKey)
		roles, err = c.SMembers(roleKey).Result()
		if err != nil {
			log.Logger.Error("Role Cache Error: " + err.Error())
			return nil, err
		}
	}

	// 排序角色
	sort.Strings(roles)
	// 删除空字符串
	for i := 0; i < len(roles); i++ {
		if roles[i] == "" {
			roles = append(roles[:i], roles[i+1:]...)
			break
		}
	}

	return roles, nil
}

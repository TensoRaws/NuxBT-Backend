package cache

import (
	"errors"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"gorm.io/gorm"
	"sort"
	"strconv"
)

const (
	ROLE_PREFIX = "role:userID:"
)

// GetUserRolesByID 根据 userID 获取用户角色列表
func GetUserRolesByID(userID int32) ([]string, error) {
	c := cache.Cache

	var roles []string

	roles, err := db.GetUserRolesByID(userID)

	// 从Redis获取用户角色
	roleKey := ROLE_PREFIX + strconv.Itoa(int(userID))

	roles, err = c.SMembers(roleKey).Result()

	// 如果Redis中没有对应的set，查询数据库
	if err != nil {
		roles, err = db.GetUserRolesByID(userID)
		if err != nil {
			// 处理数据库查询错误
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 用户没有特殊角色
				roles = []string{}
			} else {
				// 未知错误，抛了
				log.Logger.Error("RABC Unknow DB Error: " + err.Error())
				return nil, err
			}
		}

		// 将数据库中查询到的角色保存到Redis
		err = c.SAdd(roleKey, roles).Err()
		if err != nil {
			// 处理Redis存储错误
			log.Logger.Error("RABC Unknow Cache Error: " + err.Error())
			return nil, err
		}
	}

	// 排序角色
	sort.Strings(roles)

	return roles, nil
}

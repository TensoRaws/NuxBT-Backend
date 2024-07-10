package jwt

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
)

// CreateUser 新建用户
func CreateUser(user *model.User) (err error) {
	u := query.User
	err = u.Create(user)
	return err
}

// GetUserByEmail 根据 email 获取用户
func GetUserByEmail(email string) (user *model.User, err error) {
	u := query.User
	user, err = u.Where(u.Email.Eq(email)).First()
	return user, err
}

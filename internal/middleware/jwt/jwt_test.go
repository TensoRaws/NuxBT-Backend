package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/db"
)

func TestSingToken(t *testing.T) {
	// 配置初始化
	config.Init()
	log.Init()
	db.Init()

	gmail := util.GetRandomString(10) + "@gmail.com"

	// 模拟注册（新建一个用户）
	u := query.User
	user := &model.User{
		Username:   gmail,
		Email:      gmail,
		Password:   "abc",
		Signature:  "野兽先辈114514",
		Avatar:     "avatar",
		LastActive: time.Now(),
	}
	err := u.Create(user)
	if err != nil {
		t.Fatalf("create user failed, err: %v", err)
	}

	// 模拟登录（生成 JWT token）
	token := jwt.GenerateToken(gmail)
	fmt.Printf("token: %#v\n", token)

	useClaims, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("jwt token generate success, but vaild failed....")
	}

	fmt.Printf("token 解析成功，获取到的用户信息：%#v\n", *useClaims)
}

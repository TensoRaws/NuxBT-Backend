package jwt_test

import (
	"fmt"
	"testing"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/db"
)

func TestSingToken1(t *testing.T) {
	// 配置初始化
	config.Init()
	log.Init()
	db.Init()

	userName := util.GetRandomString(10) + "@gmail.com"

	// 模拟注册（新建一个用户）
	u := query.User
	user := &model.User{Name: userName, Password: "abc", Signature: "newtess", Avatar: "avatar", BackgroundImage: "background"}
	err := u.Create(user)
	if err != nil {
		t.Fatalf("create users failed, err: %v", err)
	}

	// 模拟登录（生成 JWT token）
	token := jwt.GenerateToken(userName)
	fmt.Printf("token: %#v\n", token)

	useClaims, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("jwt token generate success, but vaild failed....")
	}

	fmt.Printf("token 解析成功，获取到的用户信息：%#v\n", *useClaims)
}

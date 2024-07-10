package jwt

import (
	"errors"
	"fmt"
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var (
	TokenExpiredDuration time.Duration
	mySigningKey         []byte
)

// GetJWTTokenExpiredDuration 根据配置文件获取 jwt 的过期时间
func GetJWTTokenExpiredDuration() time.Duration {
	if TokenExpiredDuration != 0 {
		return TokenExpiredDuration
	}
	timeout := config.GetString("jwt.timeout")
	timeoutInt, _ := strconv.Atoi(timeout)
	TokenExpiredDuration = time.Minute * time.Duration(timeoutInt)
	return TokenExpiredDuration
}

// GetJWTSigningKey 根据配置文件获取 jwt 的签名密钥
func GetJWTSigningKey() []byte {
	if len(mySigningKey) != 0 {
		return mySigningKey
	}
	mySigningKey = []byte(config.GetString("jwt.key"))
	return mySigningKey
}

// GenerateTokne 用户登录成功后，根据 username 查询到用户后生成一个 token
func GenerateToken(username string) string {
	user, err := query.User.Where(query.User.Username.Eq(username)).First()
	if err != nil {
		log.Logger.Info(err)
		return ""
	}
	return GenToken(user)
}

// GenToken 生成 jwt(json web token)
func GenToken(u *model.User) string {
	userId := strconv.FormatInt(int64(u.UserID), 10)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(GetJWTTokenExpiredDuration())),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "gopher-dance",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "token",
		ID:        userId, // jwt 中保存合法用户的 ID
	}

	// 使用指定的签名算法创建用于签名的字符串对象，使用 json 序列化和 base64Url 编码生成 jwt 的 1、2 部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 以上面生成 token 作为签名值，使用 secret 进行签名获取签名值
	// 将 token 和生成的签名值使用 '.' 拼接后就生成了 jwt
	// 这里一定要使用字节切片
	tokenStr, err := token.SignedString(GetJWTSigningKey())
	if err != nil {
		log.Logger.Info(err)
		return ""
	}
	return tokenStr
}

// ParseToken 负责解析客户端 Header 中包含的 jwt，解析成功返回用户的 Claims（包含了用户的信息）
func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// 使用匿名函数先去查询服务器签名时使用的私钥，然后调用签名的验证算法进行验证
	// 验证通过后，将 tokenString 进行反编码并反序列化到 jwt.Token 结构体相应字段
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		log.Logger.Info(err)
	}

	// 对空接口类型值进行类型断言
	// 如果类型断言成功并且 token 的有效位为 true（ParseWithClaims 方法调用成功后会将 Vaild 设置为 true）
	if cliams, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return cliams, nil
	}

	return nil, errors.New("invalid token")
}

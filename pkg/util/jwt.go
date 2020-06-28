package util

import (
	"gin.example/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	jwt.StandardClaims
}

/* 生成Token */
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()

	// 过期时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username:       username,
		Password:       password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: 	"gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err

}

func ParseToken(token string) (*Claims, error) {
	// 解析鉴权的声明，内部主要是解码与校验，返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// Valid 验证基于时间的声明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


package app

import (
	"blog/global"
	"blog/internal/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func GetSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成token
func CreateToken(u *model.User) (string, error) {
	now := time.Now().Unix()
	t := time.Unix(now+global.JWTSetting.Expire, 0)
	claims := Claims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Unix(),
			Issuer:    global.JWTSetting.Isssuser,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(global.JWTSetting.Secret)
	token, err := tokenClaims.SignedString(GetSecret())
	fmt.Println(token, "token", err)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return GetSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 获取token 中的user
func GetTokenUser(c *gin.Context) *model.User {
	b := c.MustGet("user").(*model.User)
	return b
}

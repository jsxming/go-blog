package middleware

import (
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// 验证token是否正确
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			eCode = errorcode.Success
		)
		token = c.GetHeader("token")
		if token == "" {
			eCode = errorcode.InvalidToken
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				eCode = errorcode.InvalidToken
			} else {
				// 把用户信息存起来
				c.Set("user", claims.User)
			}
		}

		if eCode != errorcode.Success {
			response.Error(c, eCode)
			c.Abort()
			return
		}
		c.Next()
	}
}

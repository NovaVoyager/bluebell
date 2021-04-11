package middlewares

import (
	"github.com/miaogu-go/bluebell/controller"

	"github.com/miaogu-go/bluebell/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			controller.ResponseError(c, controller.CodeTokenInvalid)
			c.Abort()
			return
		}
		userInfo, err := jwt.ParseToken(token)
		if err != nil {
			controller.ResponseError(c, controller.CodeTokenInvalid)
			c.Abort()
			return
		}
		c.Set(controller.UserInfoKey, userInfo.UserId)
		c.Next()
	}
}

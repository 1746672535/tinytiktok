package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Auth gin的中间件, 用于验证jwt
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中拿到token
		token := ctx.DefaultQuery("token", "")
		if token == "" {
			// 如果没有token
			ctx.Set("auth", false)
			ctx.Set("msg", TokenMalformed)
			ctx.Next()
			return
		}
		// 解析jwt
		user, err := ParseToken(token)
		if err != nil {
			if errors.Is(err, TokenExpired) {
				// 如果token过期
				ctx.Set("auth", false)
				ctx.Set("msg", TokenExpired)
				ctx.Next()
				return
			}
			// 如果令牌错误
			ctx.Set("auth", false)
			ctx.Set("msg", TokenInvalid)
			ctx.Next()
			return
		}
		// 令牌有效
		ctx.Set("auth", true)
		ctx.Set("msg", TokenValid)
		ctx.Set("user", user)
		ctx.Set("userName", user.Name)
		ctx.Set("userID", user.ID)
		ctx.Next()
		return
	}
}

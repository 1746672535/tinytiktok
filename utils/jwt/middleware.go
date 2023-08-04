package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Auth gin的中间件, 用于验证jwt
func Auth() gin.HandlerFunc {
	// TODO 请将 msg 汇总
	return func(ctx *gin.Context) {
		// 从请求头中拿到token
		token := ctx.DefaultQuery("token", "")
		if token == "" {
			// 如果没有token
			ctx.Set("auth", false)
			ctx.Set("msg", "未授权")
			ctx.Next()
			return
		}
		// 解析jwt
		user, err := ParseToken(token)
		if err != nil {
			if errors.Is(err, TokenExpired) {
				// 如果token过期
				ctx.Set("auth", false)
				ctx.Set("msg", "授权过期")
				ctx.Next()
				return
			}
			// 如果没有token
			ctx.Set("auth", false)
			ctx.Set("msg", "非法授权")
			ctx.Next()
			return
		}
		ctx.Set("auth", true)
		ctx.Set("user", user)
		ctx.Set("userName", user.Name)
		ctx.Set("userId", user.ID)
		ctx.Next()
		return
	}
}

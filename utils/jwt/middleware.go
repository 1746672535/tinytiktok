package jwt

import (
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
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
			if err == TokenExpired {
				if err == TokenExpired {
					// 如果没有token
					ctx.Set("auth", false)
					ctx.Set("msg", "授权过期")
					ctx.Next()
					return
				}
			}
			// 如果没有token
			ctx.Set("auth", false)
			ctx.Set("msg", "未授权")
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

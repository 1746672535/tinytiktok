package web

import (
	"github.com/gin-gonic/gin"
	"tinytiktok/common"
	"tinytiktok/user/srv"
	"tinytiktok/utils/msg"
)

func MessageChat(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}
	srv.MessageChat(ctx)
}

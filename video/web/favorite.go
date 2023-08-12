package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/common"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
	"tinytiktok/video/proto/favorite"
	"tinytiktok/video/proto/server"
)

func Favorite(ctx *gin.Context) {
	// 鉴权 获取用户信息 拿到用户ID
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}

	// 获取参数
	userID := ctx.GetInt64("userID")

	// 访问srv层
	md := metadata.Pairs()
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, err := client.FavoriteList(metadata.NewOutgoingContext(context.Background(), md), &favorite.FavoriteListRequest{
		UserId: userID,
	})
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"video_list":  rsp.VideoList,
	})
}

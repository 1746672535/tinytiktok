package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/utils/consul"
	"tinytiktok/video/proto/favorite"
	"tinytiktok/video/proto/server"
)

func Favorite(ctx *gin.Context) {
	// 鉴权 获取用户信息 拿到用户ID
	if !ctx.GetBool("auth") {
		return
	}
	userID := ctx.GetInt64("userID")
	md := metadata.Pairs()
	// 访问srv层
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, _ := client.FavoriteList(metadata.NewOutgoingContext(context.Background(), md), &favorite.FavoriteListRequest{
		UserId: userID,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"video_list":  rsp.VideoList,
	})
}

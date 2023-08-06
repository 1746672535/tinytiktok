package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/dfs"
	"tinytiktok/video/proto/publish"
	"tinytiktok/video/proto/server"
)

func Publish(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		return
	}
	// 如果鉴权成功, 可以从ctx里面拿到用户id
	userID := ctx.GetInt64("userID")
	// 拿到用户上传的视频文件,并将其转发到存储服务器
	file, _ := ctx.FormFile("data")
	title := ctx.DefaultPostForm("title", fmt.Sprintf("用户%d上传的视频", userID))
	url, _ := dfs.UploadFile(file, userID)
	// 将请求转发到srv层
	md := metadata.Pairs()
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, _ := client.Publish(metadata.NewOutgoingContext(context.Background(), md), &publish.PublishRequest{
		AuthorId: userID,
		Title:    title,
		PlayUrl:  url,
		// TODO 请使用三方工具截取视频封面图并将其转发到存储服务器
		CoverUrl: "example.con",
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})
}

package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/utils/consul"
	"tinytiktok/video/proto/server"
	"tinytiktok/video/proto/video"
)

func Feed(ctx *gin.Context) {
	lastTime := ctx.DefaultQuery("latest_time", "-1")
	lastTimeInt, err := strconv.Atoi(lastTime)
	if err != nil {
		return
	}
	md := metadata.Pairs()
	// 鉴权
	if ctx.GetBool("auth") {
		// TODO
		fmt.Println("用户鉴权成功")
	}
	// 访问srv层
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, _ := client.Feed(metadata.NewOutgoingContext(context.Background(), md), &video.FeedRequest{
		LatestTime: int64(lastTimeInt / 1000),
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"next_time":   0,
		"video_list":  rsp.VideoList,
	},
	)
}

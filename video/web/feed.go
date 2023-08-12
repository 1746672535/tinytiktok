package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
	"tinytiktok/video/proto/feed"
	"tinytiktok/video/proto/server"
)

func Feed(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		// TODO 无需鉴权
	}

	// 获取参数
	lastTime := ctx.DefaultQuery("latest_time", "-1")
	lastTimeInt, err := strconv.Atoi(lastTime)
	userID := ctx.GetInt64("userID")
	if err != nil {
		common.ReturnErr(ctx, msg.ParameterError)
		return
	}

	// 请求srv层
	md := metadata.Pairs()
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, err := client.Feed(metadata.NewOutgoingContext(context.Background(), md), &feed.FeedRequest{
		UserId:     userID,
		LatestTime: int64(lastTimeInt / 1000),
	})
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": msg.Success,
		"status_msg":  msg.Ok,
		"next_time":   rsp.VideoList[len(rsp.VideoList)-1].CreateTime,
		"video_list":  rsp.VideoList,
	},
	)
}

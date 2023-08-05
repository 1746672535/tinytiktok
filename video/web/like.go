package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/utils/consul"
	"tinytiktok/video/proto/like"
	"tinytiktok/video/proto/server"
)

func Like(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		return
	}
	// 2. 调用srv层给视频点赞
	userID := ctx.GetInt64("userID")
	// 创建md
	md := metadata.Pairs()
	videoID := ctx.DefaultQuery("video_id", "-1")
	videoIDInt, _ := strconv.Atoi(videoID)
	actionType := ctx.DefaultQuery("action_type", "-1")
	ActionType := 2
	if actionType == "1" {
		ActionType = 1
	}
	// 访问srv层
	// TODO 请重写该方法, 点赞应该缓存至内存或redis中并延迟提交至数据库[防止用户在短时间内频繁点赞或取消点赞]
	conn := consul.GetClientConn("video-srv")
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, _ := client.Like(metadata.NewOutgoingContext(context.Background(), md), &like.LikeRequest{
		UserId:     userID,
		VideoId:    int64(videoIDInt),
		ActionType: int32(ActionType),
	})
	// 3. 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})
}

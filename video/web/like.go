package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/utils/consul"
	"tinytiktok/video/proto/like"
	"tinytiktok/video/proto/server"
)

// TODO 请重写该方法, 点赞应该缓存至内存或redis中并延迟提交至数据库[防止用户在短时间内频繁点赞或取消点赞]
func likeSrv(md metadata.MD, userId, videoId int64, isFavorite bool) (rsp *like.LikeResponse, err error) {
	// TODO 请提取为公共方法
	service, _ := consul.Reg.FindService("video-srv")
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	defer conn.Close()
	// 获取client
	client := server.NewVideoServiceClient(conn)
	// 发送请求
	actionType := 2
	if isFavorite {
		actionType = 1
	}
	rsp, _ = client.Like(metadata.NewOutgoingContext(context.Background(), md), &like.LikeRequest{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: int32(actionType),
	})
	return rsp, err
}

func Like(ctx *gin.Context) {
	// 1. 鉴权
	token, _ := ctx.Get("auth")
	if !token.(bool) {
		// 鉴权未通过
		return
	}
	// 2. 调用srv层给视频点赞
	userId, _ := ctx.Get("userId")
	// 创建md
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	videoId := ctx.DefaultQuery("video_id", "-1")
	videoIdInt, _ := strconv.Atoi(videoId)
	actionType := ctx.DefaultQuery("action_type", "-1")
	isFavorite := false
	if actionType == "1" {
		isFavorite = true
	}
	rsp, _ := likeSrv(md, userId.(int64), int64(videoIdInt), isFavorite)
	// 3. 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})

}

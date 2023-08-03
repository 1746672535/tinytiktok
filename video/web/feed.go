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
	"tinytiktok/video/proto/server"
	"tinytiktok/video/proto/video"
)

func feedSrv(md metadata.MD, lastTime int64) (rsp *video.FeedResponse, err error) {
	// TODO 请提取为公共方法
	service, _ := consul.Reg.FindService("video-srv")
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	defer conn.Close()
	// 获取client
	client := server.NewVideoServiceClient(conn)
	// 发送请求
	rsp, _ = client.Feed(metadata.NewOutgoingContext(context.Background(), md), &video.FeedRequest{
		LatestTime: lastTime,
	})
	return rsp, err
}

func Feed(ctx *gin.Context) {
	lastTime := ctx.DefaultQuery("latest_time", "-1")
	lastTimeInt, err := strconv.Atoi(lastTime)
	if err != nil {
		return
	}
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	token := ctx.DefaultQuery("token", "")
	fmt.Println(token)
	rsp, _ := feedSrv(md, int64(lastTimeInt/1000))
	var videoList []map[string]any
	for _, v := range rsp.VideoList {
		videoList = append(videoList, map[string]any{
			"id": v.Id,
			"author": map[string]any{
				"id":               v.Author.Id,
				"name":             v.Author.Name,
				"follow_count":     v.Author.FollowCount,
				"follower_count":   v.Author.FollowerCount,
				"is_follow":        true,
				"avatar":           v.Author.Avatar,
				"background_image": v.Author.BackgroundImage,
				"signature":        v.Author.Signature,
				"total_favorited":  v.Author.TotalFavorited,
				"work_count":       v.Author.WorkCount,
				"favorite_count":   v.Author.FavoriteCount,
			},
			"play_url":       v.PlayUrl,
			"cover_url":      v.CoverUrl,
			"favorite_count": v.FavoriteCount,
			"comment_count":  v.CommentCount,
			"is_favorite":    v.IsFavorite,
			"title":          v.Title,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"next_time":   0,
		"video_list":  videoList,
	},
	)
}

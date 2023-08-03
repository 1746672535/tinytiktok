package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
)

func infoSrv(md metadata.MD, userId int64) (rsp *info2.UserResponse, err error) {
	// TODO 请提取为公共方法
	service, _ := consul.Reg.FindService("user-srv")
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	defer conn.Close()
	// 获取client
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, _ = client.Info(metadata.NewOutgoingContext(context.Background(), md), &info2.UserRequest{
		UserId: userId,
	})
	return rsp, err
}

func UserInfo(ctx *gin.Context) {
	// 需要验证jwt是否通过
	auth, exist := ctx.Get("auth")
	msg, _ := ctx.Get("msg")
	if exist && !auth.(bool) {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  msg,
			"user":        map[string]any{},
		})
	}
	id := ctx.DefaultQuery("user_id", "-1")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	// 一些数据
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	// 向srv层发送请求
	rsp, _ := infoSrv(md, int64(userId))
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user": map[string]any{
			"id":               rsp.User.Id,
			"name":             rsp.User.Name,
			"follow_count":     rsp.User.FollowCount,
			"follower_count":   rsp.User.FollowerCount,
			"is_follow":        rsp.User.IsFollow,
			"avatar":           rsp.User.Avatar,
			"background_image": rsp.User.BackgroundImage,
			"signature":        rsp.User.Signature,
			"total_favorited":  rsp.User.TotalFavorited,
			"work_count":       rsp.User.WorkCount,
			"favorite_count":   rsp.User.FavoriteCount,
		},
	})
}

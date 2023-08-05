package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
)

func UserInfo(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  ctx.GetString("msg"),
			"user":        map[string]any{},
		})
		return
	}
	id := ctx.DefaultQuery("user_id", "-1")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	// 一些数据
	md := metadata.Pairs()
	// 向srv层发送请求
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, _ := client.Info(metadata.NewOutgoingContext(context.Background(), md), &info2.UserRequest{
		UserId: int64(userId),
	})
	//
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user":        rsp.User,
	})
}

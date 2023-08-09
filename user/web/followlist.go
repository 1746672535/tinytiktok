package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/user/proto/followlist"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
)

// 获取关注列表
func FollowList(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		fmt.Println("鉴权失败")
		return
	}
	// 如果鉴权成功, 可以从ctx里面拿到用户id
	userID := ctx.GetInt64("userID")
	md := metadata.Pairs()
	//访问srv层
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.FollowList(metadata.NewOutgoingContext(context.Background(), md), &followlist.FollowListRequest{
		UserId: userID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_list":   rsp.UserList,
	})

}

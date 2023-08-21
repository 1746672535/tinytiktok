package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/common"
	"tinytiktok/user/proto/friendList"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

func FriendList(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}

	userID := ctx.GetInt64("userID")
	md := metadata.Pairs()
	// 访问srv层
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.FriendList(metadata.NewOutgoingContext(context.Background(), md), &friendList.FriendListRequest{
		UserId: userID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_list":   rsp.UserList,
	})
}

package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/user/proto/followerList"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

// FollowerList 获取关注列表
func FollowerList(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}
	// 如果鉴权成功, 可以从ctx里面拿到用户id
	id := ctx.DefaultQuery("user_id", "-1")
	userID, err := strconv.Atoi(id)
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
		return
	}
	md := metadata.Pairs()
	// 访问srv层
	conn := consul.GetClientConn(common.UserServer, int64(userID))
	if conn == nil {
		panic(msg.ServerFindError)
	}
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.FollowerList(metadata.NewOutgoingContext(context.Background(), md), &followerList.FollowerListRequest{
		UserId: int64(userID),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_list":   rsp.UserList,
	})
}

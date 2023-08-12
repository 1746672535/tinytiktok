package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/user/proto/followlist"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

// FollowList 获取关注列表
func FollowList(ctx *gin.Context) {
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
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.FollowList(metadata.NewOutgoingContext(context.Background(), md), &followlist.FollowListRequest{
		UserId: int64(userID),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_list":   rsp.UserList,
	})

}

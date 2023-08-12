package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

func UserInfo(ctx *gin.Context) {
	// 鉴权是必须的
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}

	// 获取参数
	id := ctx.DefaultQuery("user_id", "-1")
	userID, err := strconv.Atoi(id)
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
		return
	}

	// md
	md := metadata.Pairs()
	// 向srv层发送请求
	conn := consul.GetClientConn(common.UserServer)
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, err := client.Info(metadata.NewOutgoingContext(context.Background(), md), &info2.UserRequest{
		UserId: int64(userID),
	})
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user":        rsp.User,
	})
}

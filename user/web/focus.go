package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/user/proto/focus"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

func Favorite(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.ServerError)
		return
	}
	// 如果鉴权成功, 可以从ctx里面拿到自己id，对方id，和操作类型
	userId := ctx.GetInt64("userID")
	md := metadata.Pairs()
	toUserId, _ := strconv.Atoi(ctx.DefaultQuery("to_user_id", "-1"))
	actionType, _ := strconv.Atoi(ctx.DefaultQuery("action_type", "-1"))

	// 访问srv层
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()

	client := server.NewUserServiceClient(conn)
	rsp, _ := client.Favorite(metadata.NewOutgoingContext(context.Background(), md), &focus.FavoriteRequest{
		UserId:     userId,
		ToUserId:   int64(toUserId),
		ActionType: int32(actionType),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})
}

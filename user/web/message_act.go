package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/user/proto/messageAct"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

func MessageAct(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}
	userID := ctx.GetInt64("userID")
	toUserId := ctx.Query("to_user_id")
	content := ctx.Query("content")
	actionType := ctx.Query("action_type")
	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("Error parsing toUserId:", err)
		ctx.JSON(http.StatusBadRequest, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Invalid toUserId"})
		return
	}
	targetActionType, err1 := strconv.ParseInt(actionType, 10, 64)
	if err1 != nil {
		log.Println("Error parsing actionType:", err)
		ctx.JSON(http.StatusBadRequest, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Invalid actionType"})
		return
	}

	md := metadata.Pairs()
	// 访问srv层
	conn := consul.GetClientConn(common.UserServer, userID)
	if conn == nil {
		panic(msg.ServerFindError)
	}
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.MessageAct(metadata.NewOutgoingContext(context.Background(), md), &messageAct.MessageActionRequest{
		UserId:     userID,
		ToUserId:   targetUserId,
		Content:    content,
		ActionType: targetActionType,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})
}

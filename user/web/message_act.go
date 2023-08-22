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
	//toUserIdInt, _ := strconv.Atoi(toUserId)
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
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, err := client.MessageAct(metadata.NewOutgoingContext(context.Background(), md), &messageAct.MessageActionRequest{
		UserId:     userID,
		ToUserId:   targetUserId,
		Content:    content,
		ActionType: targetActionType,
	})

	if err != nil {
		// 处理错误，例如打印错误日志
		log.Printf("Error calling gRPC service: %v", err)

		// 返回适当的错误响应
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": msg.Fail,
			"status_msg":  "Error calling gRPC service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
	})
}

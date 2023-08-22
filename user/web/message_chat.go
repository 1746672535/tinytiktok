package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strconv"
	"time"
	"tinytiktok/common"
	"tinytiktok/user/proto/messageChat"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
)

func MessageChat(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		common.ReturnErr(ctx, msg.AuthError)
		return
	}
	userId := ctx.GetInt64("userID")
	toUserId := ctx.Query("to_user_id")
	preMsgTime := ctx.Query("pre_msg_time")
	var latestTime time.Time
	if preMsgTime == "" {
		latestTime = time.Time{}
	} else {
		covPreMsgTime, err := strconv.ParseInt(preMsgTime, 10, 64)
		if err != nil {
			log.Println("preMsgTime 参数错误")
			ctx.JSON(http.StatusBadRequest, gin.H{"status_code": msg.Fail, "status_msg": "Invalid preMsgTime"})
			return
		}
		latestTime = time.Unix(covPreMsgTime, 0).Add(1 * time.Second)
	}
	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("toUserId 参数错误")
		ctx.JSON(http.StatusBadRequest, gin.H{"status_code": msg.Fail, "status_msg": "Invalid toUserId"})
		return
	}

	md := metadata.Pairs()
	// 访问srv层
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.MessageChat(metadata.NewOutgoingContext(context.Background(), md), &messageChat.MessageChatRequest{
		UserId:     userId,
		ToUserId:   targetUserId,
		PreMsgTime: latestTime.Unix(),
	})
	ctx.JSON(http.StatusOK, &messageChat.MessageChatResponse{
		StatusCode:  rsp.StatusCode,
		StatusMsg:   rsp.StatusMsg,
		MessageList: rsp.MessageList,
	})
}

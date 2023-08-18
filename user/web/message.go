package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"tinytiktok/user/srv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ChatResponse struct {
	Response
	MessageList []srv.Message `json:"message_list"`
}

// MessageAct 发送消息
func MessageAct(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		fmt.Println("鉴权失败")
		return
	}
	toUserId := ctx.Query("to_user_id")
	content := ctx.Query("content")
	actionType := ctx.Query("action_type")
	loginUserId := ctx.GetInt64("userID")
	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("Error parsing toUserId:", err)
		ctx.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid toUserId"})
		return
	}
	targetActionType, err1 := strconv.ParseInt(actionType, 10, 64)
	if err1 != nil {
		log.Println("Error parsing actionType:", err)
		ctx.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid actionType"})
		return
	}
	messageService := srv.GetMessageServiceInstance()
	err = messageService.SendMessage(loginUserId, targetUserId, content, targetActionType)
	if err != nil {
		log.Println("Error sending message:", err)
		ctx.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "Send Message 接口错误"})
		return
	}

	ctx.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "消息发送成功"}) // Return a success status code (200)
}

// MessageChat 消息列表
func MessageChat(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		fmt.Println("鉴权失败")
		return
	}
	loginUserId := ctx.GetInt64("userID")
	toUserId := ctx.Query("to_user_id")
	preMsgTime := ctx.Query("pre_msg_time")
	var latestTime time.Time
	if preMsgTime == "" {
		latestTime = time.Time{} // Use zero time as default
	} else {
		covPreMsgTime, err := strconv.ParseInt(preMsgTime, 10, 64)
		if err != nil {
			log.Println("preMsgTime 参数错误")
			ctx.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid preMsgTime"})
			return
		}
		latestTime = time.Unix(covPreMsgTime, 0)
	}

	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("toUserId 参数错误")
		ctx.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid toUserId"})
		return
	}
	messageService := srv.GetMessageServiceInstance()
	messages, err := messageService.MessageChat(loginUserId, targetUserId, latestTime)
	log.Println(messages)
	if err != nil {
		log.Println("Error retrieving message chat:", err)
		ctx.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0, StatusMsg: "获取消息成功"}, MessageList: messages})
	}
}

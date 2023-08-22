package srv

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/messageChat"
	"tinytiktok/utils/msg"
)

func MessageChat(ctx *gin.Context) {
	// 获取参数
	loginUserId := ctx.GetInt64("userID")
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

	// 获取消息
	messages := make([]models.ChatMessage, 0, 10)
	plainMessages, err := models.MessageChat(UserDb, loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status_code": msg.Fail, "status_msg": err.Error()})
		return
	}
	for _, tmpMessage := range plainMessages {
		var message models.ChatMessage
		message.Id = tmpMessage.Id
		message.UserId = tmpMessage.UserId
		message.ReceiverId = tmpMessage.ReceiverId
		message.MsgContent = tmpMessage.MsgContent
		message.CreatedAt = tmpMessage.CreatedAt.Unix()
		messages = append(messages, message)
	}

	// 构造响应
	var messageList []*messageChat.Message
	for _, msg := range messages {
		createTimeStr := strconv.FormatInt(msg.CreatedAt, 10)
		messageList = append(messageList, &messageChat.Message{
			Id:         msg.Id,
			ToUserId:   msg.ReceiverId,
			FromUserId: msg.UserId,
			Content:    msg.MsgContent,
			CreateTime: createTimeStr,
		})
	}
	response := &messageChat.MessageChatResponse{
		StatusCode:  msg.Success,
		StatusMsg:   "获取消息成功",
		MessageList: messageList,
	}
	ctx.JSON(http.StatusOK, response)
}

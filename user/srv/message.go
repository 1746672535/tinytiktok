package srv

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/messageAct"
	"tinytiktok/user/proto/messageChat"
	"tinytiktok/utils/msg"
)

type LatestMessage struct {
	Message string `json:"message"`
	MsgType int64  `json:"msg_type"`
}

type MessageServiceImpl struct{}

var (
	messageServiceImpl *MessageServiceImpl
	messageServiceOnce sync.Once
)

func GetMessageServiceInstance() *MessageServiceImpl {
	messageServiceOnce.Do(func() {
		messageServiceImpl = &MessageServiceImpl{}
	})
	return messageServiceImpl
}

func (messageService *MessageServiceImpl) SendMessage(fromUserId int64, toUserId int64, content string, actionType int64) error {
	var err error
	switch actionType {
	case 1:
		err = models.SendMessage(UserDb, fromUserId, toUserId, content, actionType)
	default:
		log.Println(fmt.Sprintf("未定义 actionType=%d", actionType))
		return errors.New(fmt.Sprintf("未定义 actionType=%d", actionType))
	}
	return err
}

func (messageService *MessageServiceImpl) MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]models.ChatMessage, error) {
	messages := make([]models.ChatMessage, 0, 10)
	plainMessages, err := models.MessageChat(UserDb, loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		return nil, err
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
	return messages, nil
}

func (messageService *MessageServiceImpl) LatestMessage(loginUserId int64, targetUserId int64) (LatestMessage, error) {
	plainMessage, err := models.LatestMessage(UserDb, loginUserId, targetUserId)
	if err != nil {
		log.Println("LatestMessage Service:", err)
		return LatestMessage{}, err
	}
	var latestMessage LatestMessage
	latestMessage.Message = plainMessage.MsgContent
	if plainMessage.UserId == loginUserId {
		latestMessage.MsgType = 1
	} else {
		latestMessage.MsgType = 0
	}
	return latestMessage, nil
}

func (messageService *MessageServiceImpl) SendMessageRequest(ctx *gin.Context) {
	//获取参数
	fromUserId := ctx.GetInt64("userID")
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
	err = messageService.SendMessage(fromUserId, targetUserId, content, targetActionType)
	if err != nil {
		log.Println("Error sending message:", err)
		ctx.JSON(http.StatusInternalServerError, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Send Message 接口错误"})
		return
	}
	ctx.JSON(http.StatusOK, messageAct.MessageActionResponse{StatusCode: msg.Success, StatusMsg: "消息发送成功"})
}

func (messageService *MessageServiceImpl) MessageChatRequest(ctx *gin.Context) {
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
	messages, err := messageService.MessageChat(loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("Error retrieving message chat:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status_code": msg.Fail, "status_msg": err.Error()})
		return
	}
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

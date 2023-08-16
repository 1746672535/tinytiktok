package srv

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	"tinytiktok/user/models"
)

type Message struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"from_user_id"`
	ReceiverId int64  `json:"to_user_id"`
	MsgContent string `json:"content"`
	CreatedAt  int64  `json:"create_time"`
}

type LatestMessage struct {
	message string `json:"message"`
	msgType int64  `json:"msg_type"`
}

const (
	VIDEO_INIT_NUM_PER_AUTHOR = 10
)

type MessageServiceImpl struct {
}

var (
	messageServiceImpl *MessageServiceImpl
	messageServiceOnce sync.Once
)

type MessageService interface {
	// SendMessage 发送消息服务
	SendMessage(fromUserId int64, toUserId int64, content string, actionType int64) error

	// MessageChat 聊天记录服务，注意返回的 Message 结构体字段与 Dao 层的不完全相同
	MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error)

	// LatestMessage 返回两个 loginUserId 和好友 targetUserId 最近的一条聊天记录
	LatestMessage(loginUserId int64, targetUserId int64) (LatestMessage, error)
}

// GetMessageServiceInstance Go 单例模式
func GetMessageServiceInstance() *MessageServiceImpl {
	messageServiceOnce.Do(func() {
		messageServiceImpl = &MessageServiceImpl{}
	})
	return messageServiceImpl
}

func (messageService *MessageServiceImpl) SendMessage(fromUserId int64, toUserId int64, content string, actionType int64) error {
	var err error
	switch actionType {
	// actionType = 1 发送消息
	case 1:
		err = models.SendMessage(fromUserId, toUserId, content, actionType)
	default:
		log.Println(fmt.Sprintf("未定义 actionType=%d", actionType))
		return errors.New(fmt.Sprintf("未定义 actionType=%d", actionType))
	}
	// 返回信息
	return err
}

func (messageService *MessageServiceImpl) MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error) {
	messages := make([]Message, 0, VIDEO_INIT_NUM_PER_AUTHOR)
	plainMessages, err := models.MessageChat(loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		return nil, err
	}
	err = messageService.getRespMessage(&messages, &plainMessages)
	if err != nil {
		log.Println("getRespMessage:", err)
		return nil, err
	}
	return messages, nil
}

func (messageService *MessageServiceImpl) LatestMessage(loginUserId int64, targetUserId int64) (LatestMessage, error) {
	plainMessage, err := models.LatestMessage(loginUserId, targetUserId)
	if err != nil {
		log.Println("LatestMessage Service:", err)
		return LatestMessage{}, err
	}
	var latestMessage LatestMessage
	latestMessage.message = plainMessage.MsgContent
	if plainMessage.UserId == loginUserId {
		// 最新一条消息是当前登录用户发送的
		latestMessage.msgType = 1
	} else {
		// 最新一条消息是当前好友发送的
		latestMessage.msgType = 0
	}
	return latestMessage, nil
}

// 返回 message list 接口所需的 Message 结构体
func (messageService *MessageServiceImpl) getRespMessage(messages *[]Message, plainMessages *[]models.Message) error {
	for _, tmpMessage := range *plainMessages {
		var message Message
		messageService.combineMessage(&message, &tmpMessage)
		*messages = append(*messages, message)
	}
	return nil
}

func (messageService *MessageServiceImpl) combineMessage(message *Message, plainMessage *models.Message) error {
	message.Id = plainMessage.Id
	message.UserId = plainMessage.UserId
	message.ReceiverId = plainMessage.ReceiverId
	message.MsgContent = plainMessage.MsgContent
	message.CreatedAt = plainMessage.CreatedAt.Unix()
	return nil
}

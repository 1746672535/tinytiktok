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
	Message string `json:"message"`
	MsgType int64  `json:"msg_type"`
}

var (
	messageServiceImpl *MessageServiceImpl
	messageServiceOnce sync.Once
)

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

func (messageService *MessageServiceImpl) combineMessage(message *Message, plainMessage *models.Message) {
	message.Id = plainMessage.Id
	message.UserId = plainMessage.UserId
	message.ReceiverId = plainMessage.ReceiverId
	message.MsgContent = plainMessage.MsgContent
	message.CreatedAt = plainMessage.CreatedAt.Unix()
}

type MessageServiceImpl struct {
	processedMessageIDs map[int64]bool
}

func (messageService *MessageServiceImpl) MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error) {
	messages := make([]Message, 0, 10)
	plainMessages, err := models.MessageChat(UserDb, loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		return nil, err
	}
	for _, tmpMessage := range plainMessages {
		// Check if the message ID is already processed
		if _, exists := messageService.processedMessageIDs[tmpMessage.Id]; !exists {
			var message Message
			messageService.combineMessage(&message, &tmpMessage)
			messages = append(messages, message)
			// Mark the message ID as processed
			messageService.processedMessageIDs[tmpMessage.Id] = true
		}
	}
	return messages, nil
}

func GetMessageServiceInstance() *MessageServiceImpl {
	messageServiceOnce.Do(func() {
		messageServiceImpl = &MessageServiceImpl{
			processedMessageIDs: make(map[int64]bool),
		}
	})
	return messageServiceImpl
}

package srv

import (
	"context"
	"log"
	"strconv"
	"time"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/messageChat"
	"tinytiktok/utils/msg"
)

func (h *Handle) MessageChat(ctx context.Context, req *messageChat.MessageChatRequest) (rsp *messageChat.MessageChatResponse, err error) {
	rsp = &messageChat.MessageChatResponse{}
	messages := make([]models.ChatMessage, 0, 10)
	plainMessages, err := models.MessageChat(UserDb, req.UserId, req.ToUserId, time.Unix(req.PreMsgTime, 0))
	if err != nil {
		log.Println("MessageChat Service:", err)
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = msg.Fail
		return rsp, nil
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
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success
	rsp.MessageList = messageList
	return rsp, nil
}

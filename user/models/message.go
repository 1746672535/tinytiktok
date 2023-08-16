package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	VIDEO_NUM_PER_REFRESH = 6
)

type Message struct {
	gorm.Model
	Id         int64  `json:"id" gorm:"user_id"`
	UserId     int64  `json:"user_id" gorm:"user_id"`
	ReceiverId int64  `json:"receiver_id" gorm:"receiver_id"`
	ActionType int64  `json:"action_type" gorm:"action_type"`
	MsgContent string `json:"msg_content" gorm:"msg_content"`
}

// APIMessage 返回提取的消息（仅需 Message 的部分字段）
type APIMessage struct {
	Id         int64     `json:"id"`
	MsgContent int64     `json:"content" gorm:"msg_content"`
	CreatedAt  time.Time `json:"create_time"`
}

func (Message) TableName() string {
	return "message"
}

func SaveMessage(db *gorm.DB, msg Message) error {
	result := db.Save(&msg)
	if result.Error != nil {
		log.Println("数据库保存消息失败！", result.Error)
		return result.Error
	}
	return nil
}

// SendMessage fromUserId 发送消息 content 给 toUserId
func SendMessage(db *gorm.DB, fromUserId int64, toUserId int64, content string, actionType int64) error {
	var message Message
	message.UserId = fromUserId
	message.ReceiverId = toUserId
	message.ActionType = actionType
	message.MsgContent = content
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	return SaveMessage(db, message)
}

// MessageChat 当前登录用户和其他指定用户的聊天记录
func MessageChat(db *gorm.DB, loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error) {
	messages := make([]Message, 0, VIDEO_NUM_PER_REFRESH)
	result := db.Where("(created_at > ? and created_at < ? ) and ((user_id = ? and receiver_id = ?) or (user_id = ? and receiver_id = ?))", latestTime, time.Now(), loginUserId, targetUserId, targetUserId, loginUserId).
		Order("created_at asc").
		Find(&messages)
	if result.RowsAffected == 0 {
		return messages, nil
	}
	if result.Error != nil {
		log.Println("获取聊天记录失败！")
		return nil, result.Error
	}
	return messages, nil
}

// LatestMessage 返回 loginUserId 和 targetUserId 最近的一条聊天记录
func LatestMessage(db *gorm.DB, loginUserId int64, targetUserId int64) (Message, error) {
	var message Message
	result := db.Where(&Message{UserId: loginUserId, ReceiverId: targetUserId}).
		Or(&Message{UserId: targetUserId, ReceiverId: loginUserId}).
		Order("created_at desc").Limit(1).Take(&message)
	if result.Error != nil {
		log.Println("获取最近一条聊天记录失败！")
		return Message{}, result.Error
	}
	return message, nil
}

package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Message struct {
	gorm.Model
	Id         int64  `json:"id" gorm:"user_id"`
	UserId     int64  `json:"user_id" gorm:"user_id"`
	ReceiverId int64  `json:"receiver_id" gorm:"receiver_id"`
	ActionType int64  `json:"action_type" gorm:"action_type"`
	MsgContent string `json:"msg_content" gorm:"msg_content"`
}

func (Message) TableName() string {
	return "message"
}

type ChatMessage struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"from_user_id"`
	ReceiverId int64  `json:"to_user_id"`
	MsgContent string `json:"content"`
	CreatedAt  int64  `json:"create_time"`
}

func SendMessage(db *gorm.DB, fromUserId int64, toUserId int64, content string, actionType int64) error {
	message := Message{
		UserId:     fromUserId,
		ReceiverId: toUserId,
		ActionType: actionType,
		MsgContent: content,
	}
	result := db.Save(&message)
	if result.Error != nil {
		log.Println("数据库保存消息失败！", result.Error)
		return result.Error
	}
	return nil
}

func MessageChat(db *gorm.DB, loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error) {
	message := make([]Message, 0, 6)
	result := db.Where("(created_at > ? and created_at < ? ) and ((user_id = ? and receiver_id = ?) or (user_id = ? and receiver_id = ?))", latestTime, time.Now(), loginUserId, targetUserId, targetUserId, loginUserId).
		Order("created_at asc").
		Find(&message)
	if result.RowsAffected == 0 {
		return message, nil
	}
	if result.Error != nil {
		log.Println("获取聊天记录失败！")
		return nil, result.Error
	}
	return message, nil
}

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

package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type ChatMessage struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"from_user_id"`
	ReceiverId int64  `json:"to_user_id"`
	MsgContent string `json:"content"`
	CreatedAt  int64  `json:"create_time"`
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

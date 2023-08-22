package models

import (
	"gorm.io/gorm"
	"log"
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

func MessageAct(db *gorm.DB, UserId int64, ToUserId int64, ActionType int64, Content string) error {
	message := Message{
		UserId:     UserId,
		ReceiverId: ToUserId,
		ActionType: ActionType,
		MsgContent: Content,
	}
	result := db.Save(&message)
	if result.Error != nil {
		log.Println("数据库保存消息失败！", result.Error)
		return result.Error
	}
	return nil
}

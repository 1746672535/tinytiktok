package models

import (
	"gorm.io/gorm"
	"log"
)

type LatestMessages struct {
	Message string `json:"message"`
	MsgType int64  `json:"msg_type"`
}

func LatestMessage(db *gorm.DB, loginUserId int64, targetUserId int64) (LatestMessages, error) {
	var message Message
	result := db.Where(&Message{UserId: loginUserId, ReceiverId: targetUserId}).
		Or(&Message{UserId: targetUserId, ReceiverId: loginUserId}).
		Order("created_at desc").Limit(1).Take(&message)
	if result.Error != nil {
		log.Println("获取最近一条聊天记录失败！")
		return LatestMessages{}, result.Error
	}
	return LatestMessages{
		Message: message.MsgContent,
		MsgType: 0,
	}, nil
}

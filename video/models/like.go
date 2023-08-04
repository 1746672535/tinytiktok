package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID      int64 `gorm:"primaryKey"`
	UserID  int64 `gorm:"column:user_id"`  // 点赞的用户ID
	VideoID int64 `gorm:"column:video_id"` // 被点赞的视频ID
	State   bool  `gorm:"column:state"`    // 视频是否被点赞
}

func (l Like) TableName() string {
	return "likes"
}

package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      int64  `gorm:"primaryKey" json:"id"`
	UserID  int64  `gorm:"column:user_id" json:"user_id"`
	VideoID int64  `gorm:"column:video_id" json:"video_id"`
	Content string `gorm:"column:content" json:"content"`
}

func (c Comment) TableName() string {
	return "comments"
}

// CalcCommentCountByVideoID 如果用户发表/删除评论, 则给视频的评论数量 +1 /-1  有待商榷是否将删除和增加评论操作共用到同一个接口中
func CalcCommentCountByVideoID(db *gorm.DB, videoID int64, isCommented bool) error {
	var video Video
	result := db.First(&video, videoID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("video with ID %d not found", videoID)
		} else {
			return result.Error
		}
	}
	if isCommented {
		video.CommentCount++
	} else {
		video.CommentCount--
	}
	result = db.Save(&video)
	if result.Error != nil {
		db.Rollback()
		return result.Error
	}
	return nil
}

// CommentVideo 将评论插入数据库
func CommentVideo(db *gorm.DB, comment *Comment) error {
	result := db.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCommentListByVideoID(db *gorm.DB, videoID int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteComment 用户删除视频
func DeleteComment(db *gorm.DB, commentID int64) error {
	result := db.Delete(&commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

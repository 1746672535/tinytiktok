package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey" json:"id"`
	UserID    int64  `gorm:"column:user_id" json:"user_id"`
	VideoID   int64  `gorm:"column:video_id" json:"video_id"`
	Content   string `gorm:"column:content" json:"content"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

// 评论缓存数据
type CommentCache struct {
	CommentID int64
	UserID    int64
	VideoID   int64
	Content   string
	CreatedAt string
	IsEdit    bool
}

func (c Comment) TableName() string {
	return "comments"
}

// CalcCommentCountByVideoID 如果用户发表/删除评论, 则给视频的评论数量 +1 /-1  有待商榷是否将删除和增加评论操作共用到同一个接口中
func CalcCommentCountByVideoID(db *gorm.DB, videoID int64, isCommented bool) error {
	var video Video
	result := db.First(&video, videoID)
	if result.Error != nil {
		return result.Error
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
func CommentVideo(db *gorm.DB, comment *Comment) (int64, error) {
	result := db.Create(comment)
	if result.Error != nil {
		return -1, result.Error
	}
	//将评论数量+1
	err := CalcCommentCountByVideoID(db, comment.VideoID, true)
	if err != nil {
		db.Rollback()
		return -1, err //这里到底要返回什么
	}
	var CommentId int64
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &CommentId)

	return CommentId, nil
}

// GetCommentListByVideoID 通过VideoID获取评论列表
func GetCommentListByVideoID(db *gorm.DB, videoID int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetVideoCommentsCount 获取视频评论数量
func GetVideoCommentsCount(db *gorm.DB, videoID int64) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("video_id = ? ", videoID).Count(&count)
	if result.Error != nil {
		// 查询出错
		return 0, result.Error
	}
	return count, nil
}

// DeleteComment 用户删除视频
func DeleteComment(db *gorm.DB, commentID int64) error {
	comment := Comment{}
	db.Where("id = ?", commentID).Take(&comment)
	result := db.Delete(&comment)
	//将评论数量-1
	err := CalcFavoriteCountByVideoID(db, comment.VideoID, false)
	if err != nil {
		db.Rollback()
		return err
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCommentByID 获取评论id
func GetCommentByID(db *gorm.DB, CommentID int64) (*Comment, error) {
	var c *Comment
	err := db.Where("id = ?", CommentID).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

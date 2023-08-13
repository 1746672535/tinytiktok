package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"tinytiktok/user/models"
)

type Like struct {
	gorm.Model
	ID      int64 `gorm:"primaryKey" json:"id"`
	UserID  int64 `gorm:"column:user_id" json:"user_id"`
	VideoID int64 `gorm:"column:video_id" json:"video_id"`
	State   bool  `gorm:"column:state" json:"state"`
}

func (l Like) TableName() string {
	return "likes"
}

// CalcFavoriteCountByVideoID 如果用户点赞视频, 则给视频的点赞数量 +1 / -1
func CalcFavoriteCountByVideoID(db *gorm.DB, videoID int64, isFavorite bool) error {
	var video Video
	result := db.First(&video, videoID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("video with ID %d not found", videoID)
		} else {
			return result.Error
		}
	}
	if isFavorite {
		video.FavoriteCount++
	} else {
		video.FavoriteCount--
	}
	result = db.Save(&video)
	if result.Error != nil {
		db.Rollback()
		return result.Error
	}
	return nil
}

// IsUserLikedVideo 判断用户是否点赞了指定视频
func IsUserLikedVideo(db *gorm.DB, videoID, userID int64) (bool, error) {
	var like Like
	// 查询用户是否点赞了视频
	result := db.Where("video_id = ? AND user_id = ?", videoID, userID).First(&like)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 记录未找到，表示用户未点赞该视频
			return false, nil
		}
		// 查询出错
		return false, result.Error
	}
	// 找到记录，表示用户已点赞该视频
	if like.State {
		return like.State, nil
	}
	return false, nil
}

// LikeVideo 用户点赞或取消点赞视频
func LikeVideo(db *gorm.DB, videoID, userID int64, isFavorite bool) error {
	var like Like
	result := db.Where("video_id = ? AND user_id = ?", videoID, userID).First(&like)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询出错
			return result.Error
		}
		// 记录不存在，创建新记录
		like = Like{
			UserID:  userID,
			VideoID: videoID,
			State:   isFavorite,
		}
		result = db.Create(&like)
		if result.Error != nil {
			// 添加记录出错
			return result.Error
		}
	} else {
		if like.State && isFavorite {
			// 避免重复点赞
			return errors.New("repeat likes")
		}
		// 记录已存在，更新结果和更新时间
		like.UpdatedAt = time.Now()
		like.State = isFavorite
		result = db.Save(&like)
		if result.Error != nil {
			// 如果出错则将数据回滚
			db.Rollback()
			// 更新记录出错
			return result.Error
		}
	}
	// 为视频点赞数量 +1 / -1
	err := CalcFavoriteCountByVideoID(db, videoID, isFavorite)
	if err != nil {
		db.Rollback()
		return err
	}
	// 为用户的点赞数量 +1 / -1
	CalcFavoriteCountByUserID(db, userID, isFavorite)
	return nil
}

// GetVideoLikesCount 获取视频点赞数量
func GetVideoLikesCount(db *gorm.DB, videoID int64) (int64, error) {
	var count int64
	result := db.Model(&Like{}).Where("video_id = ? AND state = 1", videoID).Count(&count)
	if result.Error != nil {
		// 查询出错
		return 0, result.Error
	}
	return count, nil
}

// IsFavorite 查询该用户是否被作者关注
func IsFavorite(db *gorm.DB, userId, authorId int64) bool {
	var user models.Relation
	result := db.Where("userid=? and pid=? and state =1 ", userId, authorId).First(&user)
	if result.Error != nil {
		return false
	}
	// 没找到
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

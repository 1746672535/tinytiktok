package srv

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/like"
)

// IsUserLikedVideo 判断用户是否点赞了指定视频
func IsUserLikedVideo(videoID, userID int64) (bool, error) {
	var like models.Like
	// 查询用户是否点赞了视频
	result := VideoDb.Where("video_id = ? AND user_id = ?", videoID, userID).First(&like)
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
		return true, nil
	}
	return false, nil
}

// LikeVideo 用户点赞或取消点赞视频
func LikeVideo(videoID, userID int64, isFavorite bool) error {
	var like models.Like
	result := VideoDb.Where("video_id = ? AND user_id = ?", videoID, userID).First(&like)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询出错
			return result.Error
		}
		// 记录不存在，创建新记录
		like = models.Like{
			UserID:  userID,
			VideoID: videoID,
			State:   isFavorite,
		}
		result = VideoDb.Create(&like)
		if result.Error != nil {
			// 添加记录出错
			return result.Error
		}
	} else {
		// 记录已存在，更新结果和更新时间
		like.UpdatedAt = time.Now()
		like.State = isFavorite
		result = VideoDb.Save(&like)
		if result.Error != nil {
			// 更新记录出错
			return result.Error
		}
	}
	return nil
}

// GetVideoLikesCount 获取视频点赞数量
func GetVideoLikesCount(videoID int64) (int64, error) {
	var count int64
	result := VideoDb.Model(&models.Like{}).Where("video_id = ? AND state = ?", videoID, true).Count(&count)
	if result.Error != nil {
		// 查询出错
		return 0, result.Error
	}
	return count, nil
}

func (h *Handle) Like(ctx context.Context, req *like.LikeRequest) (rsp *like.LikeResponse, err error) {
	rsp = &like.LikeResponse{}
	// 1 : 点赞  2 : 取消点赞
	isFavorite := req.ActionType == 1
	err = LikeVideo(req.VideoId, req.UserId, isFavorite)
	if err != nil {
		return nil, err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

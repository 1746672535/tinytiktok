package models

import (
	"context"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
	"time"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
)

type Video struct {
	gorm.Model
	ID            int64  `gorm:"primaryKey" json:"id"`
	AuthorID      int64  `gorm:"column:author_id" json:"author_id"`
	PlayURL       string `gorm:"column:play_url" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`
	Title         string `gorm:"column:title" json:"title"`
}

func (v Video) TableName() string {
	return "videos"
}

// GetUserInfo 根据id查找用户
func GetUserInfo(userId int64) (user *info2.User, err error) {
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, _ := client.Info(metadata.NewOutgoingContext(context.Background(), nil), &info2.UserRequest{
		UserId: userId,
	})
	if err != nil || rsp.StatusCode != 0 {
		return nil, err
	}
	return rsp.User, nil
}

// GetVideoList 返回数据库中的视频列表 返回截止日期之前的最多30条视频信息
func GetVideoList(db *gorm.DB, lastTime int64) []*Video {
	var videos []*Video
	db.Where("created_at < ?", time.Unix(lastTime, 0)).Order("created_at DESC").Limit(30).Find(&videos)
	return videos
}

// GetVideoById 根据视频的ID返回视频信息
func GetVideoById(db *gorm.DB, videoId int64) (*Video, error) {
	var video *Video
	result := db.Where("id = ?", videoId).First(&video)
	// 检查是否找到了对应的视频
	if result.Error != nil {
		return nil, result.Error
	}
	return video, nil
}

// GetUserFavoriteVideoList 查询用户喜欢的视频列表 1. 拿到视频的id列表 2. 根据视频id拿到视频的信息并返回
func GetUserFavoriteVideoList(db *gorm.DB, userId int64) ([]*Video, error) {
	var likes []Like
	db.Where("user_id = ? AND state = ?", userId, 1).Find(&likes)
	var videoList []*Video
	for _, like := range likes {
		video, err := GetVideoById(db, like.VideoID)
		if err != nil {
			continue
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

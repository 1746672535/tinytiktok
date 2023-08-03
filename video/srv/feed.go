package srv

import (
	"context"
	"time"
	userModels "tinytiktok/user/models"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/srv"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/video"
)

// TODO 请使用微服务交互的方式实现该功能, 此处仅为测试
// 1. 根据id查找用户 - 请使用
func getUserInfo(userId int64) (*userModels.User, error) {
	return srv.GetUserInfo(userId)
}

// GetVideoList 2. 返回数据库中的视频列表 返回截止日期之前的最多30条视频信息
func GetVideoList(lastTime int64) []*models.Video {
	var videos []*models.Video
	VideoDb.Where("created_at < ?", time.Unix(lastTime, 0)).Order("created_at DESC").Limit(30).Find(&videos)
	return videos
}

func (h *Handle) Feed(ctx context.Context, req *video.FeedRequest) (rsp *video.FeedResponse, err error) {
	rsp = &video.FeedResponse{}
	// 1. 取数据
	videos := GetVideoList(req.LatestTime)
	var videoList []*video.Video
	for _, v := range videos {
		author, err := getUserInfo(v.AuthorId)
		if err != nil {
			continue
		}
		videoList = append(videoList, &video.Video{
			Id: v.ID,
			Author: &info2.User{
				Id:              author.ID,
				Name:            author.Name,
				FollowCount:     author.FollowCount,
				FollowerCount:   author.FollowerCount,
				IsFollow:        false,
				Avatar:          author.Avatar,
				BackgroundImage: author.BackgroundImg,
				Signature:       author.Signature,
				TotalFavorited:  author.TotalFavorited,
				WorkCount:       author.WorkCount,
				FavoriteCount:   author.FavoriteCount,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			// TODO
			IsFavorite: false,
			Title:      v.Title,
		})
	}
	// 2. 返回数据
	rsp.StatusMsg = "ok"
	rsp.StatusCode = 0
	rsp.VideoList = videoList
	return rsp, nil
}

package srv

import (
	"context"
	"time"
	"tinytiktok/user/proto/info2"
	"tinytiktok/user/web"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/video"
)

// 根据id查找用户
func getUserInfo(userId int64) (user *info2.User, err error) {
	rsp, err := web.InfoSrv(nil, userId)
	if err != nil || rsp.StatusCode != 0 {
		return nil, err
	}
	return rsp.User, nil
}

// GetVideoList 返回数据库中的视频列表 返回截止日期之前的最多30条视频信息
func GetVideoList(lastTime int64) []*models.Video {
	var videos []*models.Video
	VideoDb.Where("created_at < ?", time.Unix(lastTime, 0)).Order("created_at DESC").Limit(30).Find(&videos)
	return videos
}

// Feed 获取视频列表
func (h *Handle) Feed(ctx context.Context, req *video.FeedRequest) (rsp *video.FeedResponse, err error) {
	rsp = &video.FeedResponse{}
	// 获取最近30条视频信息
	videos := GetVideoList(req.LatestTime)
	var videoList []*video.Video
	for _, v := range videos {
		// 查询视频作者信息
		author, err := getUserInfo(v.AuthorId)
		if err != nil {
			continue
		}
		// 查询该视频是否被该用户点赞
		like, err := IsUserLikedVideo(v.ID, v.AuthorId)
		if err != nil {
			continue
		}
		// 查询视频的点赞数量
		favoriteCount, err := GetVideoLikesCount(v.ID)
		if err != nil {
			continue
		}
		videoList = append(videoList, &video.Video{
			Id: v.ID,
			// 视频作者
			Author: &info2.User{
				Id:            author.Id,
				Name:          author.Name,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				// TODO 用户是否关注该作者
				IsFollow:        false,
				Avatar:          author.Avatar,
				BackgroundImage: author.BackgroundImage,
				Signature:       author.Signature,
				TotalFavorited:  author.TotalFavorited,
				WorkCount:       author.WorkCount,
				FavoriteCount:   author.FavoriteCount,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
			// 视频的点赞总数
			FavoriteCount: favoriteCount,
			// TODO 视频的评论总数
			CommentCount: v.CommentCount,
			// 该用户是否点赞
			IsFavorite: like,
			Title:      v.Title,
		})
	}
	// 2. 返回数据
	rsp.StatusMsg = "ok"
	rsp.StatusCode = 0
	rsp.VideoList = videoList
	return rsp, nil
}

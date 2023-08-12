package srv

import (
	"context"
	"tinytiktok/user/srv"
	"tinytiktok/utils/msg"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/feed"
	"tinytiktok/video/proto/video"
)

// Feed 获取视频列表
func (h *Handle) Feed(ctx context.Context, req *feed.FeedRequest) (rsp *feed.FeedResponse, err error) {
	rsp = &feed.FeedResponse{}
	// 获取最近30条视频信息
	videos := models.GetVideoList(VideoDb, req.LatestTime)
	var videoList []*video.Video
	for _, v := range videos {
		// 查询视频作者信息
		author, err := models.GetUserInfo(v.AuthorID)
		if err != nil {
			continue
		}
		// 查询该用户是否被作者关注
		if req.UserId == author.Id {
			author.IsFollow = true
		} else {
			author.IsFollow = models.IsFavorite(srv.RelationDb, req.UserId, author.Id)
		}

		// 查询该视频是否被该用户点赞
		like := false
		if req.UserId != 0 {
			like, err = models.IsUserLikedVideo(VideoDb, v.ID, req.UserId)
			if err != nil {
				continue
			}
		}
		// 查询视频的点赞数量
		favoriteCount, err := models.GetVideoLikesCount(VideoDb, v.ID)
		if err != nil {
			continue
		}
		videoList = append(videoList, &video.Video{
			Id: v.ID,
			// 视频作者
			Author:   author,
			PlayUrl:  v.PlayURL,
			CoverUrl: v.CoverURL,
			// 视频的点赞总数
			FavoriteCount: favoriteCount,
			// TODO 视频的评论总数
			CommentCount: v.CommentCount,
			// 该用户是否点赞
			IsFavorite: like,
			Title:      v.Title,
			// 视频的返回时间
			CreateTime: v.CreatedAt.UnixMilli(),
		})
	}
	// 2. 返回数据
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success
	rsp.VideoList = videoList
	return rsp, nil
}

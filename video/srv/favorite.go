package srv

import (
	"context"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/favorite"
	"tinytiktok/video/proto/video"
)

func (h *Handle) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (rsp *favorite.FavoriteListResponse, err error) {
	rsp = &favorite.FavoriteListResponse{}
	videoList, _ := models.GetUserFavoriteVideoList(VideoDb, req.UserId)
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	var videos []*video.Video
	for _, v := range videoList {
		// 查询视频作者信息
		author, err := models.GetUserInfo(v.AuthorID)
		if err != nil {
			continue
		}
		videos = append(videos, &video.Video{
			Id:            v.ID,
			Author:        author,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			// TODO
			IsFavorite: false,
			Title:      v.Title,
		})
	}
	rsp.VideoList = videos
	return rsp, nil
}

package srv

import (
	"context"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/publish"
	"tinytiktok/video/proto/video"
)

func (h *Handle) Publish(ctx context.Context, req *publish.PublishRequest) (rsp *publish.PublishResponse, err error) {
	rsp = &publish.PublishResponse{}
	video := &models.Video{
		AuthorID: req.AuthorId,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:    req.Title,
	}
	err = models.InsertVideo(VideoDb, video)
	// 为用户的作品数量+1
	err = models.CalcWorkCountByUserID(req.AuthorId, true)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "not ok"
		return rsp, err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

func (h *Handle) PublishList(ctx context.Context, req *publish.PublishListRequest) (rsp *publish.PublishListResponse, err error) {
	rsp = &publish.PublishListResponse{}
	videos, err := models.GetVideoListByUserID(VideoDb, req.UserId)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "not ok"
		return rsp, err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	var videoList []*video.Video
	for _, v := range videos {
		// 查询视频作者信息
		author, err := models.GetUserInfo(v.AuthorID)
		if err != nil {
			continue
		}
		// 查询该视频是否被该用户点赞
		like, err := models.IsUserLikedVideo(VideoDb, v.ID, req.UserId)
		if err != nil {
			continue
		}
		videoList = append(videoList, &video.Video{
			Id:            v.ID,
			Author:        author,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    like,
			Title:         v.Title,
		})
	}
	rsp.VideoList = videoList
	return rsp, nil
}

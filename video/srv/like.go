package srv

import (
	"context"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/like"
)

func (h *Handle) Like(ctx context.Context, req *like.LikeRequest) (rsp *like.LikeResponse, err error) {
	rsp = &like.LikeResponse{}
	// 1 : 点赞  2 : 取消点赞
	isFavorite := req.ActionType == 1
	err = models.LikeVideo(VideoDb, req.VideoId, req.UserId, isFavorite)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "请勿重复点赞"
		return rsp, err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

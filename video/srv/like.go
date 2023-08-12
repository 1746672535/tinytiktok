package srv

import (
	"context"
	"tinytiktok/utils/msg"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/like"
)

func (h *Handle) Like(ctx context.Context, req *like.LikeRequest) (rsp *like.LikeResponse, err error) {
	rsp = &like.LikeResponse{}
	// 1 : 点赞  2 : 取消点赞
	isFavorite := req.ActionType == 1
	err = models.LikeVideo(VideoDb, req.VideoId, req.UserId, isFavorite)
	if err != nil {
		rsp.StatusCode = msg.Fail
		rsp.StatusMsg = msg.RepeatError
		return rsp, err
	}
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}

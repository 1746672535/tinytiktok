package srv

import (
	"context"
	"tinytiktok/utils/msg"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/detail"
)

func (h *Handle) Detail(ctx context.Context, req *detail.DetailRequest) (rsp *detail.DetailResponse, err error) {
	rsp = &detail.DetailResponse{}
	d, err := models.GetDetailById(VideoDb, req.UserId)
	if err != nil {
		rsp.StatusCode = msg.Fail
		rsp.StatusMsg = msg.RepeatError
		return rsp, err
	}
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	rsp.Detail = &detail.Detail{
		UserId:         req.UserId,
		WorkCount:      d.WorkCount,
		FavoriteCount:  d.FavoriteCount,
		TotalFavorited: d.TotalFavorited,
	}
	return rsp, nil
}

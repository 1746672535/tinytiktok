package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/favorite"
	"tinytiktok/utils/msg"
)

func (h *Handle) CalcFavoriteCount(ctx context.Context, req *favorite.CalcFavoriteCountRequest) (rsp *favorite.CalcFavoriteCountResponse, err error) {
	err = models.CalcFavoriteCountByUserID(UserDb, req.UserId, req.IsFavorite)
	if err != nil {
		return nil, err
	}
	rsp = &favorite.CalcFavoriteCountResponse{
		StatusCode: msg.Success,
		StatusMsg:  msg.Ok,
	}
	return rsp, nil
}

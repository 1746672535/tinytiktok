package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/favorite"
)

func (h *Handle) CalcFavoriteCount(ctx context.Context, req *favorite.CalcFavoriteCountRequest) (rsp *favorite.CalcFavoriteCountResponse, err error) {
	err = models.CalcFavoriteCountByUserID(UserDb, req.UserId, req.IsFavorite)
	if err != nil {
		return nil, err
	}
	rsp = &favorite.CalcFavoriteCountResponse{}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

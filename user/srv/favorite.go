package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/favorite"
)

func (h *Handle) CalcFavoriteCount(ctx context.Context, req *favorite.FavoriteRequest) (rsp *favorite.FavoriteResponse, err error) {
	err = models.CalcFavoriteCountByUserID(UserDb, req.UserId, req.IsFavorite)
	if err != nil {
		return nil, err
	}
	rsp = &favorite.FavoriteResponse{}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

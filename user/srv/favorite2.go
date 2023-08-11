package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/favorite2"
)

func (h *Handle) Favorite(ctx context.Context, req *favorite2.FavoriteRequest) (rsp *favorite2.FavoriteResponse, err error) {
	rsp = &favorite2.FavoriteResponse{}
	// 1 : 关注  2 : 取消关注
	isFavorite := req.ActionType == 1
	err = models.FavoriteAction(RelationDb, req.UserId, req.ToUserId, isFavorite)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return rsp, nil
	}
	// 操作成功，对粉丝数和关注数修改
	models.UpdateCount(UserDb, req.UserId, req.ToUserId, isFavorite)
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}
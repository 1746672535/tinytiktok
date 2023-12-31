package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/focus"
	"tinytiktok/utils/msg"
)

// Favorite 关注/取消关注 操作
func (h *Handle) Favorite(ctx context.Context, req *focus.FavoriteRequest) (rsp *focus.FavoriteResponse, err error) {
	rsp = &focus.FavoriteResponse{}
	// 1 : 关注  2 : 取消关注
	isFavorite := req.ActionType == 1
	err = models.FavoriteAction(UserDb, req.UserId, req.ToUserId, isFavorite)
	if err != nil {
		rsp.StatusCode = msg.Fail
		rsp.StatusMsg = err.Error()
		return rsp, nil
	}
	// 操作成功，对粉丝数和关注数修改
	models.UpdateCount(UserDb, req.UserId, req.ToUserId, isFavorite)
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}

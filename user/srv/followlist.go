package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/followlist"
	"tinytiktok/user/proto/info2"
)

func (h *Handle) FollowList(ctx context.Context, req *followlist.FollowListRequest) (rsp *followlist.FollowListResponse, err error) {
	rsp = &followlist.FollowListResponse{}
	users := models.GetFollowList(RelationDb, req.UserId)
	var userList []*info2.User
	for _, v := range users {
		user, err := models.GetUserInfo(UserDb, v.PID)
		if err != nil {
			continue
		}
		userList = append(userList, &info2.User{
			Id:              user.ID,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImg,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		})
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	rsp.UserList = userList
	return rsp, nil
}
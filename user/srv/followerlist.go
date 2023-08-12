package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/followerlist"
	"tinytiktok/user/proto/info2"
	"tinytiktok/utils/msg"
)

func (h *Handle) FollowerList(ctx context.Context, req *followerlist.FollowerListRequest) (rsp *followerlist.FollowerListResponse, err error) {
	rsp = &followerlist.FollowerListResponse{}
	users := models.GetFollowerList(UserDb, req.UserId)
	var userList []*info2.User
	for _, v := range users {
		user, err := models.GetUserInfoF(UserDb, v.UserID)
		if err != nil {
			continue
		}
		State := models.GetStateById(UserDb, req.UserId, user.ID)
		userList = append(userList, &info2.User{
			Id:              user.ID,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        State,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImg,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		})
	}
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	rsp.UserList = userList
	return rsp, nil
}

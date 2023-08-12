package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/friendlist"
	"tinytiktok/user/proto/info2"
	"tinytiktok/utils/msg"
)

func (h *Handle) FriendList(ctx context.Context, req *friendlist.FriendListRequest) (*friendlist.FriendListResponse, error) {
	rsp := friendlist.FriendListResponse{}
	users := models.GetFriendList(UserDb, req.UserId)
	var friendList []*friendlist.FriendUser
	for _, v := range users {
		friendList = append(friendList, &friendlist.FriendUser{
			User: &info2.User{
				Id:              v.ID,
				Name:            v.Name,
				FollowCount:     v.FollowCount,
				FollowerCount:   v.FollowerCount,
				IsFollow:        v.IsFollow,
				Avatar:          v.Avatar,
				BackgroundImage: v.BackgroundImg,
				Signature:       v.Signature,
				TotalFavorited:  v.TotalFavorited,
				WorkCount:       v.WorkCount,
				FavoriteCount:   v.FavoriteCount,
			},
			Message: "开发中",
			MsgType: 0,
		})
	}
	rsp.UserList = friendList
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success

	return &rsp, nil
}

package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/followList"
	"tinytiktok/user/proto/info2"
	"tinytiktok/utils/msg"
)

func (h *Handle) FollowList(ctx context.Context, req *followList.FollowListRequest) (rsp *followList.FollowListResponse, err error) {
	rsp = &followList.FollowListResponse{}
	// 获取用户的关注列表
	users := models.GetFollowList(UserDb, req.UserId)
	var userList []*info2.User
	for _, v := range users {
		// 获取用户关注列表的用户的具体信息
		user, err := models.GetUserInfo(UserDb, v.PID)
		if err != nil {
			continue
		}
		// State := models.GetStateById(RelationDb, req.UserId, user.ID)
		userList = append(userList, &info2.User{
			Id:              user.ID,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        true,
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

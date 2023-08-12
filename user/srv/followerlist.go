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
	// 获取该用户的粉丝列表
	users := models.GetFollowerList(UserDb, req.UserId)
	var userList []*info2.User
	for _, v := range users {
		// 获取该用户粉丝列表的具体信息并查看是否关注了自己
		user, err := models.GetUserInfoF(UserDb, req.UserId, v.UserID)
		if err != nil {
			continue
		}
		// 判断是否相互关注
		State := models.IsMutualFollow(UserDb, req.UserId, user.ID)
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

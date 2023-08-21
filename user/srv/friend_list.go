package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/friendList"
	"tinytiktok/utils/msg"
)

func (h *Handle) FriendList(ctx context.Context, req *friendList.FriendListRequest) (*friendList.FriendListResponse, error) {
	rsp := friendList.FriendListResponse{}
	users := models.GetFriendList(UserDb, req.UserId)
	var friendUsers []*friendList.FriendUser
	messageService := GetMessageServiceInstance()
	for _, v := range users {
		messageInfo, err := messageService.LatestMessage(req.UserId, v.ID)
		if err != nil {
			if err.Error() != "record not found" {
				continue
			}
			messageInfo.Message = ""
		}
		friendUsers = append(friendUsers, &friendList.FriendUser{
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
			Message:         &messageInfo.Message,
			MsgType:         messageInfo.MsgType,
		})
	}
	rsp.UserList = friendUsers
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success
	return &rsp, nil
}

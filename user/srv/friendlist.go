package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/friendlist"
	"tinytiktok/utils/msg"
)

func (h *Handle) FriendList(ctx context.Context, req *friendlist.FriendListRequest) (*friendlist.FriendListResponse, error) {
	rsp := friendlist.FriendListResponse{}
	users := models.GetFriendList(UserDb, req.UserId)
	var friendList []*friendlist.FriendUser
	messageService := GetMessageServiceInstance()
	for _, v := range users {
		messageInfo, err := messageService.LatestMessage(req.UserId, v.ID)
		if err != nil {
			if err.Error() != "record not found" {
				continue
			}
			messageInfo.Message = ""
		}
		friendList = append(friendList, &friendlist.FriendUser{
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
	rsp.UserList = friendList
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success
	return &rsp, nil
}

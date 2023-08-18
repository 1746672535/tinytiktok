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

	// Get the MessageService instance
	messageService := GetMessageServiceInstance()

	for _, v := range users {
		// // Get the latest message info for the user
		messageInfo, err := messageService.LatestMessage(req.UserId, v.ID)
		// In case of an error, skip and continue to the next user
		if err != nil {
			if err.Error() != "record not found" {
				continue
			}
			messageInfo.message = ""
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
			Message:         &messageInfo.message,
			MsgType:         messageInfo.msgType,
		})
	}
	rsp.UserList = friendList
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success

	return &rsp, nil
}

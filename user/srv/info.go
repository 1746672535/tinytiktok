package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/info2"
)

// Info 用户登录
func (h *Handle) Info(ctx context.Context, req *info2.UserRequest) (rsp *info2.UserResponse, err error) {
	user, err := models.GetUserInfo(UserDb, req.UserId)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &info2.UserResponse{}
	rsp.User = &info2.User{
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
		// TODO 数据库暂未配置
		IsFollow: true,
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

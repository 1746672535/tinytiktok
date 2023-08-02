package main

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/info2"
)

func GetUserInfo(userId int64) (*models.User, error) {
	// 根据用户 Name 查询用户
	var user models.User
	result := UserDb.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 将用户信息返回
	return &user, nil
}

// Info 用户登录
func (h *Handle) Info(ctx context.Context, req *info2.UserRequest) (rsp *info2.UserResponse, err error) {
	user, err := GetUserInfo(req.UserId)
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

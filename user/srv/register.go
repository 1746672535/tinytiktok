package main

import (
	"context"
	"fmt"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/register"
	"tinytiktok/utils/tools"
)

// InsertUser 向数据库添加用户的数据
func InsertUser(username, password string) (int64, error) {
	// 生成 md5 和 salt
	encodePsd, salt := tools.GenerateMd5(password)
	// 插入数据
	user := &models.User{
		Name:           username,
		Password:       encodePsd,
		Salt:           salt,
		FollowCount:    0,
		FollowerCount:  0,
		Avatar:         "https://example.com/avatar.png",
		BackgroundImg:  "https://example.com/background.png",
		Signature:      fmt.Sprintf("Name: %s", username),
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	}
	result := UserDb.Create(user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.ID, nil
}

// Register 用户注册
func (h *Handle) Register(ctx context.Context, req *register.RegisterRequest) (rsp *register.RegisterResponse, err error) {
	userId, err := InsertUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &register.RegisterResponse{}
	rsp.UserId = userId
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

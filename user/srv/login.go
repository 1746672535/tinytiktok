package main

import (
	"context"
	"errors"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/login"
	"tinytiktok/utils/tools"
)

// LoginVerify 登录验证
func LoginVerify(username, password string) (int64, error) {
	// 根据用户 Name 查询用户
	var user models.User
	result := UserDb.Where("name = ?", username).First(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	// 验证密码是否正确
	if tools.VerifyMd5(user.Password, password, user.Salt) {
		// 登录成功
		return user.ID, nil
	}
	return -1, errors.New("密码或用户名错误")
}

// Login 用户登录
func (h *Handle) Login(ctx context.Context, req *login.LoginRequest) (rsp *login.LoginResponse, err error) {
	userId, err := LoginVerify(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &login.LoginResponse{}
	rsp.UserId = userId
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

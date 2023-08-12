package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/login"
	"tinytiktok/utils/msg"
)

// Login 用户登录
func (h *Handle) Login(ctx context.Context, req *login.LoginRequest) (rsp *login.LoginResponse, err error) {
	userID, err := models.LoginVerify(UserDb, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &login.LoginResponse{
		StatusCode: msg.Success,
		StatusMsg:  msg.Ok,
		UserId:     userID,
	}
	return rsp, nil
}

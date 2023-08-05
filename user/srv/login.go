package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/login"
)

// Login 用户登录
func (h *Handle) Login(ctx context.Context, req *login.LoginRequest) (rsp *login.LoginResponse, err error) {
	userID, err := models.LoginVerify(UserDb, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &login.LoginResponse{}
	rsp.UserId = userID
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

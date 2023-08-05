package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/register"
)

// Register 用户注册
func (h *Handle) Register(ctx context.Context, req *register.RegisterRequest) (rsp *register.RegisterResponse, err error) {
	userID, err := models.InsertUser(UserDb, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &register.RegisterResponse{}
	rsp.UserId = userID
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}

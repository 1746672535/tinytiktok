package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/register"
	"tinytiktok/utils/msg"
)

// Register 用户注册
func (h *Handle) Register(ctx context.Context, req *register.RegisterRequest) (rsp *register.RegisterResponse, err error) {
	userID, err := models.InsertUser(UserDb, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 返回信息
	rsp = &register.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  msg.Ok,
		UserId:     userID,
	}
	return rsp, nil
}

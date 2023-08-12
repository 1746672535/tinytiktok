package srv

import (
	"context"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/publish"
	"tinytiktok/utils/msg"
)

func (h *Handle) CalcWorkCount(ctx context.Context, req *publish.CalcWorkCountRequest) (rsp *publish.CalcWorkCountResponse, err error) {
	err = models.CalcWorkCountByUserID(UserDb, req.UserId, req.IsPublish)
	if err != nil {
		return nil, err
	}
	rsp = &publish.CalcWorkCountResponse{}
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}

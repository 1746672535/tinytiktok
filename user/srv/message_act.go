package srv

import (
	"context"
	"log"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/messageAct"
	"tinytiktok/utils/msg"
)

func (h *Handle) MessageAct(ctx context.Context, req *messageAct.MessageActionRequest) (rsp *messageAct.MessageActionResponse, err error) {
	rsp = &messageAct.MessageActionResponse{}
	err = models.MessageAct(UserDb, req.UserId, req.ToUserId, req.ActionType, req.Content)
	if err != nil {
		log.Println("Error sending message:", err)
		rsp.StatusMsg = "Send Message 接口错误"
		rsp.StatusCode = msg.Fail
		return rsp, nil
	}
	rsp.StatusMsg = msg.Ok
	rsp.StatusCode = msg.Success
	return rsp, nil
}

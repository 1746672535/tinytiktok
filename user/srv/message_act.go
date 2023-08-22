package srv

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/messageAct"
	"tinytiktok/utils/msg"
)

func MessageAct(ctx *gin.Context) {
	// 获取参数
	fromUserId := ctx.GetInt64("userID")
	toUserId := ctx.Query("to_user_id")
	content := ctx.Query("content")
	actionType := ctx.Query("action_type")
	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("Error parsing toUserId:", err)
		ctx.JSON(http.StatusBadRequest, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Invalid toUserId"})
		return
	}
	targetActionType, err1 := strconv.ParseInt(actionType, 10, 64)
	if err1 != nil {
		log.Println("Error parsing actionType:", err)
		ctx.JSON(http.StatusBadRequest, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Invalid actionType"})
		return
	}

	err = models.SendMessage(UserDb, fromUserId, targetUserId, content, targetActionType)
	if err != nil {
		log.Println("Error sending message:", err)
		ctx.JSON(http.StatusInternalServerError, messageAct.MessageActionResponse{StatusCode: msg.Fail, StatusMsg: "Send Message 接口错误"})
		return
	}

	ctx.JSON(http.StatusOK, messageAct.MessageActionResponse{StatusCode: msg.Success, StatusMsg: "消息发送成功"})
}

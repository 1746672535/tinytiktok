package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"tinytiktok/user/srv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ChatResponse struct {
	Response
	MessageList []srv.Message `json:"message_list"`
}

// MessageAction 发送消息
func MessageAct(c *gin.Context) {
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	actionType := c.Query("action_type")
	loginUserId := c.GetInt64("userId")

	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("Error parsing toUserId:", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid toUserId"})
		return
	}

	targetActionType, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		log.Println("Error parsing actionType:", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid actionType"})
		return
	}

	messageService := srv.GetMessageServiceInstance()
	err = messageService.SendMessage(loginUserId, targetUserId, content, targetActionType)
	if err != nil {
		log.Println("Error sending message:", err)
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "Send Message 接口错误"})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat 消息列表
func MessageChat(c *gin.Context) {
	loginUserId := c.GetInt64("userId")
	toUserId := c.Query("to_user_id")
	preMsgTime := c.Query("pre_msg_time")

	covPreMsgTime, err := strconv.ParseInt(preMsgTime, 10, 64)
	if err != nil {
		log.Println("Error parsing preMsgTime:", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid preMsgTime"})
		return
	}

	latestTime := time.Unix(covPreMsgTime, 0)

	targetUserId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		log.Println("Error parsing toUserId:", err)
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Invalid toUserId"})
		return
	}

	messageService := srv.GetMessageServiceInstance()
	messages, err := messageService.MessageChat(loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("Error retrieving message chat:", err)
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0, StatusMsg: "获取消息成功"}, MessageList: messages})
}

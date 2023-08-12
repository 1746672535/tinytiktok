package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"tinytiktok/utils/config"
	"tinytiktok/utils/msg"
)

var UserServer string
var VideoServer string

func init() {
	path := os.Getenv("APP")
	cfg := config.NewConfig(fmt.Sprintf("%s/config", path), "mysql.yaml", "yaml")
	UserServer = cfg.ReadString("User.Name")
	VideoServer = cfg.ReadString("Video.Name")
}

func ReturnErr(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": msg.Fail,
		"status_msg":  message,
	})
}

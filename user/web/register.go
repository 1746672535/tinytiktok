package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/user/proto/register"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/jwt"
)

func UserRegister(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")
	// md
	md := metadata.Pairs()
	// 向srv层发送请求
	conn := consul.GetClientConn("user-srv")
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, _ := client.Register(metadata.NewOutgoingContext(context.Background(), md), &register.RegisterRequest{
		Username: username,
		Password: password,
	})
	rsp.Token, _ = jwt.CreateToken(&jwt.UserClaims{
		ID:   rsp.UserId,
		Name: username,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_id":     rsp.UserId,
		"token":       rsp.Token,
	})
}

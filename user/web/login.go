package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/common"
	"tinytiktok/user/proto/login"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/jwt"
	"tinytiktok/utils/msg"
)

func UserLogin(ctx *gin.Context) {
	// 获取参数
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")
	if username == "" || password == "" {
		common.ReturnErr(ctx, msg.ParameterError)
		return
	}

	// md
	md := metadata.Pairs()
	// 向srv层发送请求
	conn := consul.GetClientConn(common.UserServer)
	defer conn.Close()
	client := server.NewUserServiceClient(conn)
	rsp, err := client.Login(metadata.NewOutgoingContext(context.Background(), md), &login.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
		return
	}

	// 获取jwt
	rsp.Token, err = jwt.CreateToken(&jwt.UserClaims{
		ID:   rsp.UserId,
		Name: username,
	})
	if err != nil {
		common.ReturnErr(ctx, msg.JwtError)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_id":     rsp.UserId,
		"token":       rsp.Token,
	})
}

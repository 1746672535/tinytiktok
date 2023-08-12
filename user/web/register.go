package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/common"
	"tinytiktok/user/proto/register"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/jwt"
	"tinytiktok/utils/msg"
)

func UserRegister(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")
	// 错误处理
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
	rsp, err := client.Register(metadata.NewOutgoingContext(context.Background(), md), &register.RegisterRequest{
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

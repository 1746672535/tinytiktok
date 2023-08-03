package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"tinytiktok/user/proto/login"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/jwt"
)

func loginSrv(md metadata.MD, username, password string) (rsp *login.LoginResponse, err error) {
	// TODO 请提取为公共方法
	service, _ := consul.Reg.FindService("user-srv")
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	defer conn.Close()
	// 获取client
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, _ = client.Login(metadata.NewOutgoingContext(context.Background(), md), &login.LoginRequest{
		Username: username,
		Password: password,
	})
	return rsp, err
}

func UserLogin(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")
	// 一些数据
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	// 向srv层发送请求
	rsp, _ := loginSrv(md, username, password)
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

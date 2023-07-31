package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"tinytiktok/user/proto/register"
	"tinytiktok/user/proto/server"
)

// TODO
func registerSrv(md metadata.MD, username, password string) (rsp *register.RegisterResponse, err error) {
	service, _ := reg.FindService("user-srv")
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	defer conn.Close()
	// 获取client
	client := server.NewUserServiceClient(conn)
	// 发送请求
	rsp, _ = client.Register(metadata.NewOutgoingContext(context.Background(), md), &register.RegisterRequest{
		Username: username,
		Password: password,
	})
	return rsp, err
}

// TODO
func registerUser(ctx *gin.Context) {
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	rsp, err := registerSrv(md, "Jiudan", "520666")
	fmt.Println(rsp, err)
}

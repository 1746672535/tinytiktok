package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
	"os"
	"tinytiktok/user/proto/server"
	"tinytiktok/user/srv"
	"tinytiktok/utils/config"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/tools"
)

func main() {
	// 初始化配置文件
	path := os.Getenv("APP")
	serverConfig := config.NewConfig(fmt.Sprintf("%s\\config", path), "server.yaml", "yaml")
	// 启动服务
	g := grpc.NewServer()
	server.RegisterUserServiceServer(g, &srv.Handle{})
	// 启用注册中心
	id := uuid.NewString()
	ip := tools.GetLocalIP()
	port := tools.GetFreePort()
	listen, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// 注册服务
	reg := consul.NewRegistry(serverConfig.ReadString("Consul.Host"), serverConfig.ReadInt("Consul.Port"))
	reg.Register(&consul.Server{
		Address: ip,
		Port:    port,
		Name:    serverConfig.ReadString("User.Name"),
		Id:      id,
		Tags:    serverConfig.ReadStringSlice("User.Tag"),
		HealthCheck: consul.HealthCheck{
			TCP:                            fmt.Sprintf("%s:%d", ip, port),
			Timeout:                        serverConfig.ReadString("User.Timeout"),
			Interval:                       serverConfig.ReadString("User.Interval"),
			DeregisterCriticalServiceAfter: serverConfig.ReadString("User.DeregisterCriticalServiceAfter"),
		},
	})
	// 延迟注销服务
	defer reg.DeRegister(id)
	_ = g.Serve(listen)
}

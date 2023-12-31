package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
	"os"
	"tinytiktok/utils/config"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
	"tinytiktok/utils/tools"
	"tinytiktok/video/proto/server"
	"tinytiktok/video/srv"
)

func main() {
	// 初始化配置文件
	path := os.Getenv("APP")
	serverConfig := config.NewConfig(fmt.Sprintf("%s/config", path), "server.yaml", "yaml")
	// 启动服务
	g := grpc.NewServer()
	server.RegisterVideoServiceServer(g, &srv.Handle{})
	// 启用注册中心
	id := uuid.NewString()
	ip := tools.GetLocalIP()
	port := tools.GetFreePort()
	listen, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// 注册服务
	reg := consul.NewRegistry(serverConfig.ReadString("Consul.Host"), serverConfig.ReadInt("Consul.Port"))
	err := reg.Register(&consul.Server{
		Address: ip,
		Port:    port,
		Name:    serverConfig.ReadString("Video.Name"),
		ID:      id,
		Tags:    serverConfig.ReadStringSlice("Video.Tag"),
		HealthCheck: consul.HealthCheck{
			TCP:                            fmt.Sprintf("%s:%d", ip, port),
			Timeout:                        serverConfig.ReadString("Video.Timeout"),
			Interval:                       serverConfig.ReadString("Video.Interval"),
			DeregisterCriticalServiceAfter: serverConfig.ReadString("Video.DeregisterCriticalServiceAfter"),
		},
	})
	if err != nil {
		panic(msg.ServerRegisterError + err.Error())
	}
	// 延迟注销服务
	defer reg.DeRegister(id)
	// 需要在服务退出之前将redis中的数据同步到mysql中
	defer srv.FlushLikeDataToMysql()
	_ = g.Serve(listen)
}

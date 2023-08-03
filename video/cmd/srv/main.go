package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/tools"
	"tinytiktok/video/proto/server"
	"tinytiktok/video/srv"
)

func main() {
	// 启动服务
	g := grpc.NewServer()
	server.RegisterVideoServiceServer(g, &srv.Handle{})
	// 启用注册中心
	id := uuid.NewString()
	ip := tools.GetLocalIP()
	port := tools.GetFreePort()
	listen, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// 注册服务
	reg := consul.NewRegistry("127.0.0.1", 8500)
	reg.Register(&consul.Server{
		Address: ip,
		Port:    port,
		Name:    "video-srv",
		Id:      id,
		Tags:    []string{"video", "srv"},
		HealthCheck: consul.HealthCheck{
			TCP:                            fmt.Sprintf("%s:%d", ip, port),
			Timeout:                        "3s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})
	// 延迟注销服务
	defer reg.DeRegister(id)
	_ = g.Serve(listen)
}

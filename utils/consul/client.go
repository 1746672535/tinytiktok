package consul

import (
	"fmt"
	"google.golang.org/grpc"
)

func GetClientConn(serverName string) *grpc.ClientConn {
	service, _ := Reg.FindService(serverName)
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	return conn
}

package consul

import (
	"fmt"
	"google.golang.org/grpc"
)

func GetClientConn(serverName string) *grpc.ClientConn {
	service, err := Reg.FindService(serverName)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return nil
	}
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
	return conn
}

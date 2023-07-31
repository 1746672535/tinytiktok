package consul

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

var reg *Registry

func init() {
	reg = NewRegistry("127.0.0.1", 8500)
}

func TestNewServer(t *testing.T) {
	server := &Server{
		Address: "127.0.0.1",
		Port:    8087,
		Name:    "tiny-tiktok-test",
		Id:      uuid.NewString(),
		Tags:    []string{"tiny", "tiktok"},
		HealthCheck: HealthCheck{
			HTTP:                           "127.0.0.1:8087",
			Timeout:                        "3s",
			Interval:                       "15s",
			DeregisterCriticalServiceAfter: "20s",
		},
	}
	err := reg.Register(server)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestDeRegister(t *testing.T) {
	server, _ := reg.FindService("demo")
	err := reg.DeRegister(server.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestFindServer(t *testing.T) {
	server, err := reg.FindService("demo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(server.Address, server.Port)
}

func TestDeRegisterAllServer(t *testing.T) {
	reg.DeRegisterAllServer()
}

package consul

import (
	"fmt"
	"github.com/awnumar/fastrand"
	"github.com/hashicorp/consul/api"
	"os"
	"os/signal"
	"syscall"
)

type HealthCheck struct {
	// 注意Consul是在docker里面运行, 指定本机地址
	TCP                            string
	HTTP                           string
	Timeout                        string
	Interval                       string
	DeregisterCriticalServiceAfter string
}

type Server struct {
	// 服务地址端口
	Address string
	Port    int
	Name    string
	Id      string
	Tags    []string
	// 健康检查
	HealthCheck HealthCheck
}

type Registry struct {
	Config api.Config
}

func NewRegistry(address string, port int) *Registry {
	r := &Registry{Config: api.Config{
		//Token:   token,
		Address: fmt.Sprintf("%s:%d", address, port),
	}}
	return r
}

// Register 注册服务
func (r *Registry) Register(server *Server) error {
	client, err := api.NewClient(&r.Config)
	if err != nil {
		return err
	}
	//生成注册对象
	registration := api.AgentServiceRegistration{
		Name:    server.Name,
		ID:      server.Id,
		Address: server.Address,
		Port:    server.Port,
		Tags:    server.Tags,
		Check: &api.AgentServiceCheck{
			TCP:                            server.HealthCheck.TCP,
			HTTP:                           server.HealthCheck.HTTP,
			Timeout:                        server.HealthCheck.Timeout,
			Interval:                       server.HealthCheck.Interval,
			DeregisterCriticalServiceAfter: server.HealthCheck.DeregisterCriticalServiceAfter,
		},
	}
	// 注册
	err = client.Agent().ServiceRegister(&registration)
	// 监听 Ctrl C 信号, 在程序中断时自动注销服务
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		_ = r.DeRegister(server.Id)
		os.Exit(0)
	}()
	return err
}

// DeRegister 注销服务
func (r *Registry) DeRegister(serverID string) error {
	client, err := api.NewClient(&r.Config)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serverID)
	return err
}

// DeRegisterAllServer 注销所有服务
func (r *Registry) DeRegisterAllServer() {
	client, _ := api.NewClient(&r.Config)
	services, _ := client.Agent().Services()
	for _, v := range services {
		_ = client.Agent().ServiceDeregister(v.ID)
	}
}

// FindService 发现服务
func (r *Registry) FindService(ServerName string) (*api.AgentService, error) {
	// 加载客户端
	client, _ := api.NewClient(&r.Config)
	// 获取服务
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%v"`, ServerName))
	if err != nil {
		return nil, err
	}
	// 返回结果
	var result []*api.AgentService
	for _, service := range data {
		result = append(result, service)
	}
	return result[fastrand.Intn(len(result))], nil
}

package consul

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/awnumar/fastrand"
	"github.com/hashicorp/consul/api"
	"math"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"tinytiktok/utils/config"
)

var Reg *Registry

func init() {
	// 初始化配置文件
	path := os.Getenv("APP")
	consulConfig := config.NewConfig(fmt.Sprintf("%s/config", path), "server.yaml", "yaml")
	host := consulConfig.ReadString("Consul.Host")
	port := consulConfig.ReadInt("Consul.Port")
	Reg = NewRegistry(host, port)
}

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
	ID      string
	Tags    []string
	// 健康检查
	HealthCheck HealthCheck
}

type ServerInfo struct {
	Address string
	Port    int
	Weight  int
}

type Registry struct {
	Config api.Config
}

func NewRegistry(address string, port int) *Registry {
	r := &Registry{Config: api.Config{
		// Token:   token,
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
	// 生成注册对象
	registration := api.AgentServiceRegistration{
		Name:    server.Name,
		ID:      server.ID,
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
		_ = r.DeRegister(server.ID)
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
		// 检查服务的状态
		checks, _, err := client.Agent().AgentHealthServiceByID(service.ID)
		if err != nil {
			return nil, errors.New("no service")
		}
		if checks != "passing" {
			continue
		}
		result = append(result, service)
	}
	if len(result) == 0 {
		return nil, errors.New("no service")
	}
	return result[fastrand.Intn(len(result))], nil
}

// FindServiceList 返回服务列表
func (r *Registry) FindServiceList(ServerName string) ([]*api.AgentService, error) {
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
		// 检查服务的状态
		checks, _, err := client.Agent().AgentHealthServiceByID(service.ID)
		if err != nil {
			return nil, errors.New("no service")
		}
		if checks != "passing" {
			continue
		}
		result = append(result, service)
	}
	if len(result) == 0 {
		return nil, errors.New("no service")
	}
	return result, nil
}

// 调度算法
// note 将服务的id和用户的id hash成一个0-1的浮点数[hash的值是固定的, 仅与id相关], 这样我们就可以使得一个用户固定的访问其中一个实例
func hashIntToFloat(value int) float64 {
	// 将整数转换为字符串
	strValue := strconv.Itoa(value)
	// 使用哈希函数计算字符串的哈希值
	const prime32 = uint32(16777619)
	hash := uint32(2166136261)
	for i := 0; i < len(strValue); i++ {
		hash ^= uint32(strValue[i])
		hash *= prime32
	}
	hashValue := hash
	// 将哈希值映射到 0 到 1 之间的浮点数范围, 这里使用的是简单的线性映射
	floatValue := float64(hashValue) / math.MaxUint32
	return floatValue
}
func hashUUIDToFloat(uuid string) float64 {
	// 将 UUID 字符串转换为字节数组
	uuidBytes := []byte(uuid)
	// 使用 SHA-256 哈希算法对字节数组进行哈希
	hashBytes := sha256.Sum256(uuidBytes)
	// 将哈希结果解释为一个大整数
	hashInt := binary.BigEndian.Uint64(hashBytes[:])
	// 将大整数映射到 0 到 1 之间的浮点数范围, 这里使用的是简单的线性映射
	floatValue := float64(hashInt) / math.MaxUint64
	return floatValue
}
func findClosestGreater(A float64, B []float64) float64 {
	closest := math.MaxFloat64
	for _, num := range B {
		if num > A && num < closest {
			closest = num
		}
	}
	return closest
}

package consul

import (
	"fmt"
	"github.com/awnumar/fastrand"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"strconv"
	"testing"
	"tinytiktok/utils/tools"
)

var r *Registry
var ServerName = "tiny-tiktok-test"

func init() {
	r = NewRegistry("127.0.0.1", 8500)
}

// 测试添加服务
func TestNewServer(t *testing.T) {
	for i := 0; i < 4; i++ {
		weight := fastrand.Intn(100) + 1
		server := &Server{
			Address: "127.0.0.1",
			Port:    tools.GetFreePort(),
			Name:    ServerName,
			ID:      uuid.NewString(),
			// 第一个参数应该是权重值, 我们建议根据服务器性能合理设置, 值应该是一个整数[1-100]
			Tags: []string{strconv.Itoa(weight), "tiny", "tiktok", "test"},
		}
		// 注册一个失败的服务
		if i == 3 {
			server.HealthCheck = HealthCheck{
				TCP:      "127.0.0.1:5788",
				Interval: "5s",
				Timeout:  "7s",
			}
		}
		_ = r.Register(server)
	}
}

// 测试注销服务
func TestDeRegister(t *testing.T) {
	server, _ := r.FindService(ServerName)
	err := r.DeRegister(server.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 测试查找服务
func TestFindServer(t *testing.T) {
	server, err := r.FindService(ServerName)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(server.Address, server.Port)
}

// 测试注销所有服务
func TestDeRegisterAllServer(t *testing.T) {
	r.DeRegisterAllServer()
}

// 测试查找所有服务
func TestFindAllServer(t *testing.T) {
	// 加载客户端
	client, _ := api.NewClient(&r.Config)
	// 获取服务
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%v"`, ServerName))
	if err != nil {
		panic(err)
	}
	// 返回结果
	for _, service := range data {
		checks, info, err := client.Agent().AgentHealthServiceByID(service.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println(checks, info.Service.Port, info.Checks.AggregatedStatus())
	}
}

// 测试hash
func TestHash(t *testing.T) {
	var floatList []float64
	for i := 0; i < 10; i++ {
		floatList = append(floatList, hashUUIDToFloat(uuid.NewString()))
	}
	uid := hashIntToFloat(1)
	fmt.Println(floatList)
	fmt.Println(uid)
	fmt.Println(findClosestGreater(uid, floatList))
}

// 测试调度算法
// note 第一步将所有健康的服务的权重拿出来, 例如 10, 30, 60 那么他们的值应该分别为0.1, 0.3, 0.6
// note 然后按照从小到大的顺序为他们分配区间 [0.0-0.10) [0.10, 0.40) [0.40, 1), 命中该区间的用户应该访问固定的服务 [我们假设所有用户是没有差异的, 无高欲望用户和低欲望用户的区别]

func TestServiceScheduling(t *testing.T) {
	// 用户idhash的值
	userHash := 0.57888
	// 第一步 获取所有可用服务
	serverList, _ := r.FindServiceList(ServerName)
	var AvailableServices []*ServerInfo
	var TotalWeight int
	for _, service := range serverList {
		weight, err := strconv.Atoi(service.Tags[0])
		if err != nil {
			continue
		}
		TotalWeight += weight
		AvailableServices = append(AvailableServices, &ServerInfo{
			Address: service.Address,
			Port:    service.Port,
			Weight:  weight,
		})
	}
	// 第二步为其分配命中区间
	start := 0.0
	for _, service := range AvailableServices {
		start += float64(service.Weight) / float64(TotalWeight)
		// 如果用户的hash值小于该值则命中
		if start > userHash {
			fmt.Println(service.Address, service.Port)
			return
		}
	}
}

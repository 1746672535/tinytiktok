package consul

import (
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// GetClientConn userID是一个可选参数
func GetClientConn(serverName string, userID ...int64) *grpc.ClientConn {
	// 如果用户ID为0, 则随机生成一个hash值
	userHash := 0.0
	if len(userID) == 0 {
		rand.Seed(time.Now().UnixNano())
		userHash = rand.Float64()
	} else {
		userHash = hashIntToFloat(int(userID[0]))
	}
	// 第一步 获取所有可用服务
	serverList, _ := Reg.FindServiceList(serverName)
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
	// 排序
	sort.Slice(AvailableServices, func(i, j int) bool {
		return AvailableServices[i].Weight < AvailableServices[j].Weight
	})
	// 第二步为其分配命中区间
	start := 0.0
	for _, service := range AvailableServices {
		start += float64(service.Weight) / float64(TotalWeight)
		// 如果用户的hash值小于该值则命中
		if start > userHash {
			conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port), grpc.WithInsecure())
			return conn
		}
	}
	return nil
}

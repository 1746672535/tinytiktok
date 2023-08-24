package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"os"
	"strconv"
	"time"
	"tinytiktok/utils/config"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var Client *redis.Client
var RefreshTime int
var ExpireTime int

func init() {
	// 初始化配置文件
	path := os.Getenv("APP")
	redisConfig := config.NewConfig(fmt.Sprintf("%s/config", path), "redis.yaml", "yaml")
	// 创建Redis客户端
	host := redisConfig.ReadString("Host")
	port := redisConfig.ReadInt("Port")
	password := redisConfig.ReadString("Password")
	RefreshTime = redisConfig.ReadInt("RefreshTime")
	ExpireTime = redisConfig.ReadInt("ExpireTime")
	if ExpireTime <= RefreshTime {
		panic("Redis Config Error: 过期时间必须大于刷新时间")
	}
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
	})
}

// Key 生成Redis的key
func Key(keys ...any) string {
	joinedStr := ""
	for _, value := range keys {
		joinedStr += fmt.Sprintf("%v-", value)
	}
	return joinedStr[:len(joinedStr)-1]
}

// Set 存储任意数据结构 - 将结构体数据转换为json
func Set(key string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if Client.Set(key, string(jsonData), time.Duration(ExpireTime)*time.Second).Err() != nil {
		return err
	}
	return nil
}

// Get 获取任意数据结构 - 将json数据转换为结构体
func Get(key string, dataStruct any) error {
	data, err := Client.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), &dataStruct)
}

// Del 删除redis的数据
func Del(key string) error {
	return Client.Del(key).Err()
}

// Exists 判断是否存在key对应的数据
func Exists(key string) (bool, error) {
	has, err := Client.Exists(key).Result()
	if has == 1 {
		return true, err
	}
	return false, err
}

// 将string类型的值转换为any
func restoreValue(value string) any {
	// 尝试将值解析为整数
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}
	// 尝试将值解析为浮点数
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return floatValue
	}
	return value
}

// 将map[string]string转换为map[string]any
func restoreValues(data map[string]string) map[string]any {
	result := make(map[string]any)
	for key, value := range data {
		// 尝试将值解析为整数
		if intValue, err := strconv.Atoi(value); err == nil {
			result[key] = intValue
			continue
		}
		// 尝试将值解析为浮点数
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			result[key] = floatValue
			continue
		}
		// 默认情况下，将值保留为字符串
		result[key] = value
	}
	return result
}

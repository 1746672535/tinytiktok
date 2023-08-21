package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // Redis服务器密码，如果有的话
		DB:       0,                // 使用的Redis数据库
	})

	// 存储Hash类型数据
	err := client.HSet("user:1", "name", "John Doe").Err()
	if err != nil {
		panic(err)
	}

	// 获取Hash类型数据的某个字段值
	name, err := client.HGet("user:1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)

	// 获取Hash类型数据的所有字段和值
	user, err := client.HGetAll("user:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("User:", user)

	// 删除Hash类型数据的某个字段
	err = client.HDel("user:1", "name").Err()
	if err != nil {
		panic(err)
	}

	// 检查Hash类型数据的某个字段是否存在
	exists, err := client.HExists("user:1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name exists:", exists)
}

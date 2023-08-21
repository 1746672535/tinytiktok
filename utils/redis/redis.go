package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"reflect"
	"strconv"
	"time"
	"tinytiktok/utils/config"
)

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

func Key(keys ...any) string {
	joinedStr := ""
	for _, value := range keys {
		joinedStr += fmt.Sprintf("%v-", value)
	}
	return joinedStr[:len(joinedStr)-1]
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

// HSet 设置key中某一个字段的值
func HSet(key, filed string, value any) error {
	return Client.HSet(key, filed, value).Err()
}

// HGet 获取key中某一个字段的值
func HGet(key, field string) (any, error) {
	data := Client.HGet(key, field)
	return restoreValue(data.Val()), data.Err()
}

// PutHash 将结构体数据存储到redis中
func PutHash(key string, obj any) error {
	data := structToMap(obj)
	err := Client.HMSet(key, data).Err()
	if err != nil {
		return err
	}
	return Client.Expire(key, time.Duration(ExpireTime)*time.Second).Err()
}

// GetHash 根据key获取结构体数据
func GetHash(key string, obj any) error {
	data := Client.HGetAll(key)
	if data.Err() != nil {
		return data.Err()
	}
	mapToStruct(obj, restoreValues(data.Val()))
	return nil
}

// 将结构体转换为map
func structToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()
	// 创建map
	result := make(map[string]interface{})
	// 遍历结构体字段
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		// 将字段名和字段值存储到map中
		result[field.Name] = fieldValue.Interface()
	}
	return result
}

// 将map转换为结构体 note 请务必传递指针类型数据, 否则将导致程序崩溃
func mapToStruct(obj any, data map[string]interface{}) any {
	objValue := reflect.ValueOf(obj)
	// 检查是否为指针类型
	if objValue.Kind() != reflect.Ptr {
		panic("请传递指针类型数据")
	}
	// 使用反射获取结构体类型
	objType := reflect.TypeOf(obj).Elem()
	// 使用反射设置结构体字段的值
	for key, value := range data {
		field, found := objType.FieldByName(key)
		if found {
			fieldValue := objValue.Elem().FieldByName(field.Name)
			if fieldValue.IsValid() && fieldValue.CanSet() {
				// 根据字段类型进行不同的处理
				switch fieldValue.Kind() {
				case reflect.Bool:
					if value == 1 {
						fieldValue.SetBool(true)
					}
				default:
					// 其他类型进行通用的转换处理
					fieldValue.Set(reflect.ValueOf(value).Convert(field.Type))
				}
			}
		}
	}
	return obj
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

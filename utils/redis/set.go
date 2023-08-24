package redis

import (
	"time"
)

// SAdd 存储集合
func SAdd[K int | int64 | float64 | string | bool | any](key string, data []K) error {
	values := make([]any, len(data))
	for i, item := range data {
		values[i] = item
	}
	if err := Client.SAdd(key, values...).Err(); err != nil {
		return err
	}
	return Client.Expire(key, time.Duration(ExpireTime)*time.Second).Err()
}

// SGet 读取集合
func SGet(key string) ([]string, error) {
	return Client.SMembers(key).Result()
}

package redis

import (
	"github.com/go-redis/redis"
	"time"
)

// ZAdd 添加有序集合 - 请在传递数据之前排序
func ZAdd[K int | int64 | float64 | string | bool | any](key string, data []K) error {
	values := make([]any, len(data))
	for i, item := range data {
		values[i] = item
	}
	var zData []redis.Z
	for i, d := range values {
		zData = append(zData, redis.Z{
			Score:  float64(i),
			Member: d,
		})
	}
	// 使用ZAdd方法将ZSET数据存储到Redis中
	if err := Client.ZAdd(key, zData...).Err(); err != nil {
		return err
	}
	return Client.Expire(key, time.Duration(ExpireTime)*time.Second).Err()
}

// ZAppend 有序集合向后追加
func ZAppend[K int | int64 | float64 | string | bool | any](key string, data []K) error {
	values := make([]any, len(data))
	for i, item := range data {
		values[i] = item
	}
	var zData []redis.Z
	scoreStart := int(Client.ZCard(key).Val())
	for i, d := range values {
		zData = append(zData, redis.Z{
			Score:  float64(scoreStart + i),
			Member: d,
		})
	}
	// 使用ZAdd方法将ZSET数据存储到Redis中 - 追加数据应该延长数据过期时间
	if err := Client.ZAdd(key, zData...).Err(); err != nil {
		return err
	}
	return Client.Expire(key, time.Duration(ExpireTime)*time.Second).Err()
}

// ZGet 读取有序集合
func ZGet(key string) ([]string, error) {
	return Client.ZRange(key, 0, -1).Result()
}

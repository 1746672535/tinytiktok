package redis

import "time"

// LPush 向Redis中添加List类型数据
func LPush[T int | int64 | float64 | string | bool | any](key string, data []T) error {
	values := make([]any, len(data))
	for i, item := range data {
		values[i] = item
	}
	if err := Client.LPush(key, values...).Err(); err != nil {
		return err
	}
	// 新增数据应该为其延迟过期时间
	return Client.Expire(key, time.Duration(ExpireTime)*time.Second).Err()
}

// LGet 获取List类型数据
func LGet(key string) ([]string, error) {
	return Client.LRange(key, 0, -1).Result()
}

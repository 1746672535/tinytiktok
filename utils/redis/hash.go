package redis

import (
	"errors"
	"reflect"
	"time"
)

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
	if len(data.Val()) == 0 {
		return errors.New("key is not exist")
	}
	mapToStruct(obj, restoreValues(data.Val()))
	return nil
}

// HSet 设置key中某一个字段的值
func HSet(key, filed string, value any) error {
	return Client.HSet(key, filed, value).Err()
}

// HGet 获取key中某一个字段的值
func HGet(key, field string) (any, error) {
	data := Client.HGet(key, field)
	if data.Err() != nil {
		return "", data.Err()
	}
	if len(data.Val()) == 0 {
		return "", errors.New("key is not exist")
	}
	return restoreValue(data.Val()), nil
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

// 将map转换为结构体
func mapToStruct(obj any, data map[string]interface{}) any {
	objValue := reflect.ValueOf(obj)
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

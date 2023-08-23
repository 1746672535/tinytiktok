package redis

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	IsStu   bool
	Country string
}

type Person2 struct {
	Name    string
	Age     int
	IsStu   bool
	Country string
}

func TestRedis(t *testing.T) {
	p1 := Person{
		Name:    "Alice",
		Age:     24,
		IsStu:   true,
		Country: "USA",
	}
	// note 尽量使用指针类型传递值
	_ = PutHash("user-1", &p1)

	// note 只要字段名一致就可以映射
	p2 := &Person2{}
	_ = GetHash("user-1", p2)
	fmt.Println(p2)

	// HSet note 请注意字段区分大小写
	_ = HSet("user-1", "Name", "Alma")

	// HGet
	value, _ := HGet("user-1", "Name")
	fmt.Println(value)

	// Exists
	exists, _ := Exists("user-1")
	fmt.Println(exists)

	// Del
	_ = Del("user-1")
}

func TestReflect(t *testing.T) {
	// 创建结构体实例
	person := Person{
		Name:    "John",
		Age:     21,
		Country: "USA",
	}

	// 将结构体转换为map[string]interface{}
	personMap := structToMap(&person)
	fmt.Println(personMap)

	// 创建一个新的map[string]interface{}
	newPersonMap := map[string]interface{}{
		"Name":    "Jack",
		"Age":     24,
		"Country": "UK",
	}

	// 将map[string]interface{}转换为结构体
	p := Person{}
	mapToStruct(&p, newPersonMap)
	fmt.Println(p)
}

func TestRedisKey(t *testing.T) {
	err := GetHash("kk", &Person{})
	if err != nil {
		fmt.Println(err)
	}
}

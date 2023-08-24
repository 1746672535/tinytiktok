package redis

import (
	"fmt"
	"testing"
)

type Address struct {
	Country   string
	Provinces string
	City      string
}

type Person struct {
	Name    string
	Age     int
	IsStu   bool
	Address Address
}

type Person2 struct {
	Name  string
	Age   int
	IsStu bool
}

func TestSetAndGet(t *testing.T) {
	_ = Set(Key("user", 1, 1), Person{
		Name:  "Alice",
		Age:   18,
		IsStu: true,
		Address: Address{
			Country:   "USA",
			Provinces: "State of New York",
			City:      "NewYork",
		},
	})
	p := Person{}
	_ = Get(Key("user", 1, 1), &p)
	fmt.Println(p.Address.Provinces)
}

func TestSetAddAndGet(t *testing.T) {
	// 存储数据
	data := []int{6, 7, 8, 9}
	_ = SAdd("test", data)
	// 读取数据
	data2, _ := SGet("test")
	fmt.Println(data2)
}

func TestZSetAddAndGet(t *testing.T) {
	// 存储数据 - 请在存储数据之前排序
	data := []any{"alice", 18, true}
	_ = ZAdd("z_test", data)
	// 读取数据
	data2, _ := ZGet("z_test")
	fmt.Println(data2)
	// 追加数据 - 请避免 true 和 1 混搅
	appendData := []int64{1, 2, 3}
	_ = ZAppend("z_test", appendData)
	// 读取数据
	data3, _ := ZGet("z_test")
	fmt.Println(data3)
}

func TestListSetAndGet(t *testing.T) {
	data := []any{"alice", 18, true}
	_ = LPush("list_test", data)
	// 获取数据
	data2, _ := LGet("list_test")
	fmt.Println(data2)
}

func TestHashSetAndGet(t *testing.T) {
	p1 := Person{
		Name:  "Alice",
		Age:   24,
		IsStu: true,
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
		Name: "John",
		Age:  21,
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

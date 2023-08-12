package yiyan

import (
	"github.com/levigross/grequests"
)

type YiYan struct {
	Content  string
	origin   string
	Author   string
	Category string
}

func GenYiYan() string {
	data, _ := grequests.Get("https://v1.jinrishici.com/all.json", nil)
	y := &YiYan{}
	data.JSON(y)
	return y.Content
}

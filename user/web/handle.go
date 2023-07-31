package main

import (
	"github.com/gin-gonic/gin"
	"tinytiktok/utils/consul"
)

var reg *consul.Registry

func init() {
	reg = consul.NewRegistry("127.0.0.1", 8500)
}

// TODO
func main() {
	router := gin.Default()
	router.POST("/douyin/user/register/", registerUser)
	router.Run(":5051")
}

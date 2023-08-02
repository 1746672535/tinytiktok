package main

import (
	"github.com/gin-gonic/gin"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/jwt"
)

var reg *consul.Registry

func init() {
	reg = consul.NewRegistry("127.0.0.1", 8500)
}

// TODO
func main() {
	router := gin.Default()
	router.Use(jwt.JwtAuth())
	router.POST("/douyin/user/register/", userRegister)
	router.POST("/douyin/user/login/", userLogin)
	router.GET("/douyin/user/", userInfo)
	router.Run(":5051")
}

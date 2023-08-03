package main

import (
	"github.com/gin-gonic/gin"
	userWeb "tinytiktok/user/web"
	"tinytiktok/utils/jwt"
	videoWeb "tinytiktok/video/web"
)

// TODO
func main() {
	router := gin.Default()
	router.Use(jwt.JwtAuth())
	// User
	router.POST("/douyin/user/register/", userWeb.UserRegister)
	router.POST("/douyin/user/login/", userWeb.UserLogin)
	router.GET("/douyin/user/", userWeb.UserInfo)
	// Video
	router.GET("/douyin/feed/", videoWeb.Feed)
	router.Run(":5051")
}

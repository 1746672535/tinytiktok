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
	router.Use(jwt.Auth())
	// User
	router.POST("/douyin/user/register/", userWeb.UserRegister)
	router.POST("/douyin/user/login/", userWeb.UserLogin)
	router.GET("/douyin/user/", userWeb.UserInfo)
	// Video
	router.GET("/douyin/feed/", videoWeb.Feed)
	router.POST("/douyin/favorite/action/", videoWeb.Like)
	router.GET("/douyin/favorite/list/", videoWeb.Favorite)
	router.POST("/douyin/publish/action/", videoWeb.Publish)
	router.GET("/douyin/publish/list/", videoWeb.PublishList)
	_ = router.Run(":5051")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userWeb "tinytiktok/user/web"
	"tinytiktok/utils/dfs"
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
	// dfs
	router.POST("/dfs/auth/", dfsAuth)
	// 关注列表
	router.GET("/douyin/relation/follow/list/", userWeb.FollowList)
	// 粉丝列表
	router.GET("/douyin/relation/follower/list/", userWeb.FollowerList)

	_ = router.Run(":5051")
}

// 用于验证 dfs-key, 防止服务器被恶意上传文件
func dfsAuth(ctx *gin.Context) {
	authToken, _ := ctx.GetPostForm("auth_token")
	if authToken != dfs.Key {
		ctx.String(http.StatusOK, "fail")
		return
	}
	ctx.String(http.StatusOK, "ok")
}

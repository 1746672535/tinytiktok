package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	userWeb "tinytiktok/user/web"
	"tinytiktok/utils/config"
	"tinytiktok/utils/dfs"
	"tinytiktok/utils/jwt"
	videoWeb "tinytiktok/video/web"
)

func main() {
	router := gin.Default()
	router.Use(jwt.Auth())
	// config
	path := os.Getenv("APP")
	cfg := config.NewConfig(fmt.Sprintf("%s/config", path), "server.yaml", "yaml")
	gin.SetMode(cfg.ReadString("Mode"))
	// dfs
	router.POST("/dfs/auth/", dfs.Auth)
	// 注册
	router.POST("/douyin/user/register/", userWeb.UserRegister)
	// 登录
	router.POST("/douyin/user/login/", userWeb.UserLogin)
	// 用户信息
	router.GET("/douyin/user/", userWeb.UserInfo)
	// 视频列表
	router.GET("/douyin/feed/", videoWeb.Feed)
	// 点赞
	router.POST("/douyin/favorite/action/", videoWeb.Like)
	// 喜欢列表
	router.GET("/douyin/favorite/list/", videoWeb.Favorite)
	// 发布视频
	router.POST("/douyin/publish/action/", videoWeb.Publish)
	// 发布列表
	router.GET("/douyin/publish/list/", videoWeb.PublishList)
	// 关注列表
	router.GET("/douyin/relation/follow/list/", userWeb.FollowList)
	// 粉丝列表
	router.GET("/douyin/relation/follower/list/", userWeb.FollowerList)
	// 关注操作
	router.POST("/douyin/relation/action/", userWeb.Favorite)
	// 好友列表
	router.GET("/douyin/relation/friend/list/", userWeb.FriendList)
	// 发送信息
	router.POST("/douyin/message/action/", userWeb.MessageAct)
	// 聊天记录
	router.GET("/douyin/message/chat/", userWeb.MessageChat)
	// 发表/删除评论
	router.POST("/douyin/comment/action/", videoWeb.Comment)
	// 评论列表
	router.GET("/douyin/comment/list/", videoWeb.CommentList)
	// 启动服务
	_ = router.Run(fmt.Sprintf(":%d", cfg.ReadInt("Server.Port")))
}

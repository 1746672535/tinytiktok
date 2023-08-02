package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 用户ID
	ID int64 `gorm:"column:id"`
	// 用户名称 - 一般是手机号
	Name string `gorm:"column:name;unique"`
	// 用户密码 - 请使用 md5 + salt
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
	// 关注总数
	FollowCount int64 `gorm:"column:follow_count"`
	// 粉丝数量
	FollowerCount int64 `gorm:"column:follower_count"`
	// 头像地址
	Avatar string `gorm:"column:avatar"`
	// 顶部大图
	BackgroundImg string `gorm:"column:background_image"`
	// 个人简介
	Signature string `gorm:"column:signature"`
	// 总获赞数量
	TotalFavorited int64 `gorm:"column:total_favorited"`
	// 作品数量
	WorkCount int64 `gorm:"column:work_count"`
	// 点赞数量
	FavoriteCount int64 `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

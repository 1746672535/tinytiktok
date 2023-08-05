package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tinytiktok/utils/tools"
)

type User struct {
	gorm.Model
	ID             int64  `gorm:"column:id" json:"id"`
	Name           string `gorm:"column:name;unique" json:"name"`
	Password       string `gorm:"column:password" json:"-"`
	Salt           string `gorm:"column:salt" json:"-"`
	FollowCount    int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount  int64  `gorm:"column:follower_count" json:"follower_count"`
	Avatar         string `gorm:"column:avatar" json:"avatar"`
	BackgroundImg  string `gorm:"column:background_image" json:"background_image"`
	Signature      string `gorm:"column:signature" json:"signature"`
	TotalFavorited int64  `gorm:"column:total_favorited" json:"total_favorited"`
	WorkCount      int64  `gorm:"column:work_count" json:"work_count"`
	FavoriteCount  int64  `gorm:"column:favorite_count" json:"favorite_count"`
}

func (User) TableName() string {
	return "users"
}

// GetUserInfo 获取用户信息
func GetUserInfo(db *gorm.DB, userID int64) (*User, error) {
	// 根据用户 Name 查询用户
	var user User
	result := db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 将用户信息返回
	return &user, nil
}

// LoginVerify 登录验证
func LoginVerify(db *gorm.DB, username, password string) (int64, error) {
	// 根据用户 Name 查询用户
	var user User
	result := db.Where("name = ?", username).First(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	// 验证密码是否正确
	if tools.VerifyMd5(user.Password, password, user.Salt) {
		// 登录成功
		return user.ID, nil
	}
	return -1, errors.New("密码或用户名错误")
}

// InsertUser 向数据库添加用户的数据
func InsertUser(db *gorm.DB, username, password string) (int64, error) {
	// 生成 md5 和 salt
	encodePsd, salt := tools.GenerateMd5(password)
	// 插入数据
	user := &User{
		Name:           username,
		Password:       encodePsd,
		Salt:           salt,
		FollowCount:    0,
		FollowerCount:  0,
		Avatar:         "https://example.com/avatar.png",
		BackgroundImg:  "https://example.com/background.png",
		Signature:      fmt.Sprintf("Name: %s", username),
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	}
	result := db.Create(user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.ID, nil
}

// CalcFavoriteCountByUserID 计算用户的喜欢数量
func CalcFavoriteCountByUserID(db *gorm.DB, userID int64, isFavorite bool) error {
	var user User
	// 从数据库中获取对应用户的信息
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}
	if isFavorite {
		user.FavoriteCount++
	} else {
		user.FavoriteCount--
	}
	// 将修改后的User结构体保存回数据库
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

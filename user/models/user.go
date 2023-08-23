package models

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
	"tinytiktok/common"
	"tinytiktok/utils/avatar"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
	"tinytiktok/utils/tools"
	"tinytiktok/utils/yiyan"
	"tinytiktok/video/proto/detail"
	"tinytiktok/video/proto/server"
)

type User struct {
	gorm.Model
	ID             int64  `gorm:"column:id" json:"id"`
	Name           string `gorm:"column:name;unique" json:"name"`
	Password       string `gorm:"column:password" json:"-"`
	Salt           string `gorm:"column:salt" json:"-"`
	FollowCount    int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount  int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow       bool   `gorm:"column:is_follow" json:"is_follow"`
	Avatar         string `gorm:"column:avatar" json:"avatar"`
	BackgroundImg  string `gorm:"column:background_image" json:"background_image"`
	Signature      string `gorm:"column:signature" json:"signature"`
	WorkCount      int64  `gorm:"-" json:"work_count"`
	FavoriteCount  int64  `gorm:"-" json:"favorite_count"`
	TotalFavorited int64  `gorm:"-" json:"total_favorited"`
}

func (User) TableName() string {
	return "users"
}

// GetUserInfo 获取用户信息
func GetUserInfo(db *gorm.DB, userID int64) (*User, error) {
	// 根据用户 Name 查询用户
	var user User
	result := db.Where("id = ?", userID).First(&user)
	d, err := GetDetail(userID)
	if result.Error != nil || err != nil {
		return nil, result.Error
	}
	user.WorkCount = d.WorkCount
	user.FavoriteCount = d.FavoriteCount
	user.TotalFavorited = d.TotalFavorited
	// 将用户信息返回
	return &user, nil
}

// GetDetail 获取用户有关视频的详细信息
func GetDetail(userID int64) (*detail.Detail, error) {
	conn := consul.GetClientConn(common.VideoServer, userID)
	if conn == nil {
		return nil, errors.New(msg.ServerFindError)
	}
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	// 发送请求
	rsp, err := client.Detail(metadata.NewOutgoingContext(context.Background(), nil), &detail.DetailRequest{
		UserId: userID,
	})
	if err != nil || rsp.StatusCode != 0 {
		return nil, err
	}
	return rsp.Detail, nil
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
		Avatar:         avatar.Gen(username),
		BackgroundImg:  "https://images.pexels.com/photos/956981/milky-way-starry-sky-night-sky-star-956981.jpeg?auto=compress&cs=tinysrgb&w=600",
		Signature:      yiyan.GenYiYan(),
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

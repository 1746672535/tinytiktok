package moudle

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             int64  `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	FollowCount    int64  `gorm:"column:follow_count"`
	FollowerCount  int64  `gorm:"column:follower_count"`
	IsFollow       bool   `gorm:"column:is_follow"`
	Avatar         string `gorm:"column:avatar"`
	BackgroundImg  string `gorm:"column:background_image"`
	Signature      string `gorm:"column:signature"`
	TotalFavorited int64  `gorm:"column:total_favorited"`
	WorkCount      int64  `gorm:"column:work_count"`
	FavoriteCount  int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

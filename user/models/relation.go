package models

import (
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	ID     int64 `gorm:"column:id" json:"id"`
	UserID int64 `gorm:"column:userid" json:"userid"` // 自己
	PID    int64 `gorm:"column:pid" json:"pid"`       // 好友
}

func (table *Relation) TableName() string {
	return "relation"
}

// 根据id查询好友列表
func GetFollowList(db *gorm.DB, userId int64) []*Relation {
	var relation []*Relation
	db.Where("userid=?", userId).Find(&relation)
	return relation
}

// 根据id查询粉丝列表
func GetFollowerList(db *gorm.DB, userId int64) []*Relation {
	var relation []*Relation
	db.Where("pid=?", userId).Find(&relation)
	return relation
}

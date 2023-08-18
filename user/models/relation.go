package models

import (
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	ID     int64 `gorm:"column:id" json:"id"`
	UserID int64 `gorm:"column:userid" json:"userid"` // 自己
	PID    int64 `gorm:"column:pid" json:"pid"`       // 好友
	State  bool  `gorm:"column:state" json:"state"`   // 状态 bool->uid关注pid  false uid没关注pid
}

func (table *Relation) TableName() string {
	return "relation"
}

// GetFollowList 根据id查询关注列表
func GetFollowList(db *gorm.DB, userId int64) []*Relation {
	var relation []*Relation
	result := db.Where("userid=? and state=1", userId).Find(&relation)
	if result.Error != nil {
		return nil
	}
	return relation
}

// GetFollowerList 根据id查询粉丝列表
func GetFollowerList(db *gorm.DB, userId int64) []*Relation {
	var relation []*Relation
	db.Where("pid=? and state =1", userId).Find(&relation)
	return relation
}

// IsMutualFollow 查看是否互相关注
func IsMutualFollow(db *gorm.DB, userId, pId int64) bool {
	var relation Relation
	result := db.Where("userid=? and pid=? and state= 1", userId, pId).Find(&relation)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// GetFriendList 获取好友列表/判断是否为好友关系
func GetFriendList(db *gorm.DB, userId int64) []*User {
	var userList []*User
	// 获取自己的关注列表好友信息
	followUsers := GetFollowList(db, userId)
	// 根据自己的关注好友列表的信息查看谁关注了自己
	for _, v := range followUsers {
		if IsMutualFollow(db, v.PID, userId) {
			user, err := GetUserInfo(db, v.PID)
			if err != nil {
				continue
			}
			user.IsFollow = true
			userList = append(userList, user)
		}
	}
	return userList
}

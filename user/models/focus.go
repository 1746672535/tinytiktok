package models

import (
	"errors"
	"gorm.io/gorm"
)

// FavoriteAction 关注操作
func FavoriteAction(db *gorm.DB, userId, toUserId int64, isFavorite bool) error {
	var relation Relation
	result := db.Where("user_id=? and pid=? ", userId, toUserId).First(&relation)
	if result.Error != nil {
		// 查询出错
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		// 记录不存在，关注-》创造新记录，取消关注-》返回错误
		if isFavorite { // 关注操作
			relation = Relation{
				UserID: userId,
				PID:    toUserId,
				State:  true,
			}
			result = db.Create(&relation)
			if result.Error != nil {
				// 添加记录出错
				return result.Error
			}
			return nil
		} else {
			return errors.New("未关注，不存在取消关注操作")
		}
	}
	// 有数据 + 关注操作
	if isFavorite {
		// 重复关注
		if relation.State {
			return errors.New("repeat favorite")
		}
		// 关注
		relation.State = true
		result = db.Save(&relation)

	} else { // 有数据 + 取消关注
		// 取消关注操作
		relation.State = false
		result = db.Save(&relation)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// UpdateCount 更新关注/粉丝数
func UpdateCount(db *gorm.DB, userId, toUserId int64, isFavorite bool) error {
	var user1 User
	var user2 User
	result1 := db.First(&user1, userId)
	if result1.Error != nil {
		return result1.Error
	}
	result2 := db.First(&user2, toUserId)
	if result2.Error != nil {
		return result2.Error
	}

	// 关注操作
	if isFavorite {
		user1.FollowCount++
		user2.FollowerCount++
	} else {
		user1.FollowCount--
		user2.FollowerCount--
	}
	result1 = db.Save(&user1)
	result2 = db.Save(&user2)

	return nil
}

package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// FavoriteAction 关注操作
func FavoriteAction(db *gorm.DB, userId, toUserId int64, isFavorite bool) error {
	var relation Relation
	result := db.Where("userid=? and pid=?", userId, toUserId).First(&relation)
	if result.Error != nil {
		// 查询出错
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		// 记录不存在，创造新记录
		if isFavorite { // 关注操作
			relation = Relation{
				UserID: userId,
				PID:    toUserId,
			}
			result = db.Create(&relation)
			if result.Error != nil {
				// 添加记录出错
				return result.Error
			}
			return nil
		}
	}

	if isFavorite {
		// 重复关注
		return errors.New("repeat favorite")
	} else {
		// 取消关注操作
		db.Delete(&relation)
		return nil
	}
}

// UpdateCount 更新关注/粉丝数
func UpdateCount(db *gorm.DB, userId, toUserId int64, isFavorite bool) error {
	var user1 User
	var user2 User
	result1 := db.First(&user1, userId)
	if result1.Error != nil {
		if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d not found", userId)
		} else {
			return result1.Error
		}
	}
	result2 := db.First(&user2, toUserId)
	if result2.Error != nil {
		if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d not found", toUserId)
		} else {
			return result2.Error
		}
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

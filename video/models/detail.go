package models

import (
	"errors"
	"gorm.io/gorm"
)

// Detail 存放用户的点赞数量, 作品数量, 被点赞数量
type Detail struct {
	gorm.Model
	ID int64 `gorm:"column:id" json:"id"`
	// 用户的作品数量
	WorkCount int64 `gorm:"column:work_count" json:"work_count"`
	// 用户喜欢的数量
	FavoriteCount int64 `gorm:"column:favorite_count" json:"favorite_count"`
	// 用户所有作品被点赞数量的总和 获赞数量
	TotalFavorited int64 `gorm:"column:total_favorited" json:"total_favorited"`
}

func (v Detail) TableName() string {
	return "details"
}

// GetDetailById 获取用户有关视频的详细信息
func GetDetailById(db *gorm.DB, userID int64) (*Detail, error) {
	var detail *Detail
	// 检查是否找到了对应的信息
	if result := db.First(&detail, userID); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询出错
			return nil, result.Error
		}
		// 记录不存在，创建新记录
		detail = &Detail{
			ID: userID,
		}
		result = db.Create(&detail)
		if result.Error != nil {
			// 添加记录出错
			return nil, result.Error
		}
	}
	return detail, nil
}

// CalcFavoriteCountByUserID 计算喜欢的数量
func CalcFavoriteCountByUserID(db *gorm.DB, userID int64, isFavorite bool) (err error) {
	var detail Detail
	// 从数据库中获取对应用户的信息
	if result := db.First(&detail, userID); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询出错
			return result.Error
		}
		// 记录不存在，创建新记录
		detail = Detail{
			ID: userID,
		}
		result = db.Create(&detail)
		if result.Error != nil {
			// 添加记录出错
			return result.Error
		}
		return err
	}
	if isFavorite {
		detail.FavoriteCount++
	} else {
		detail.FavoriteCount--
	}
	// 将修改后的User结构体保存回数据库
	if err := db.Save(&detail).Error; err != nil {
		return err
	}
	return nil
}

// CalcWorkCountByUserID 计算作品的数量
func CalcWorkCountByUserID(db *gorm.DB, userID int64, isPublish bool) (err error) {
	var detail Detail
	// 从数据库中获取对应用户的信息
	if result := db.First(&detail, userID); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 查询出错
			return result.Error
		}
		// 记录不存在，创建新记录
		detail = Detail{
			ID: userID,
		}
		result = db.Create(&detail)
		if result.Error != nil {
			// 添加记录出错
			return result.Error
		}
		return err
	}
	if isPublish {
		detail.WorkCount++
	} else {
		detail.WorkCount--
	}
	// 将修改后的User结构体保存回数据库
	if err := db.Save(&detail).Error; err != nil {
		return err
	}
	return nil
}

package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID int64 `gorm:"primaryKey"`
	// 作者的ID
	AuthorId      int64  `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
}

func (v Video) TableName() string {
	return "videos"
}

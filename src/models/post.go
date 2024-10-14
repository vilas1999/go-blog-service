package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey;autoIncrement"`
	Content string
	Title   string
	UserId  uint64 `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserId"`
}

func (Post) TableName() string {
	return "posts"
}

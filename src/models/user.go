package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	EmailId  string `gorm:"uniqueIndex;not null"`
	Password string
}

func (User) TableName() string {
	return "users"
}

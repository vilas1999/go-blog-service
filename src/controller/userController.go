package controller

import (
	"go-blog-service/src/db"
	"go-blog-service/src/models"

	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func CreateUserController() UserController {
	appDB := db.GetAppDB()
	return UserController{
		db: appDB,
	}
}

func (u *UserController) AddUser(email string, password string) (*models.User, error) {

	user := &models.User{
		EmailId:  email,
		Password: password,
	}

	res := u.db.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

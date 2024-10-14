package controller

import (
	"fmt"
	"go-blog-service/src/db"
	"go-blog-service/src/models"

	"gorm.io/gorm"
)

type PostsContoller struct {
	db *gorm.DB
}

func CreatePostsContoller() PostsContoller {
	appDB := db.GetAppDB()
	return PostsContoller{
		db: appDB,
	}
}

func (p *PostsContoller) AddPost(content string, title string, userId uint64) (*models.Post, error) {

	post := &models.Post{
		Content: content,
		Title:   title,
		UserId:  userId,
	}

	res := p.db.Create(post)

	if res.Error != nil {
		return nil, res.Error
	}

	return post, nil
}

func (p *PostsContoller) GetPost(id string) (*models.Post, error) {
	var reponsePost *models.Post
	res := p.db.Find(&reponsePost, "Id = ?", id)

	if res.Error != nil {
		fmt.Println("errror", res.Error)
		return nil, res.Error
	}

	return reponsePost, nil
}

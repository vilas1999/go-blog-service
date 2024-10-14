package db

import (
	"go-blog-service/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=127.0.0.1 user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	localDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}
	db = localDb
}

func GetAppDB() *gorm.DB {

	err := db.AutoMigrate(&models.Post{})
	err1 := db.AutoMigrate(&models.User{})
	if err != nil || err1 != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	return db
}

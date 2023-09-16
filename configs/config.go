package configs

import (
	articleDatabase "article-api/models/article/database"
	commentDatabase "article-api/models/comment/database"
	likeDatabase "article-api/models/like/database"
	profileDatabase "article-api/models/profile/database"
	userDatabase "article-api/models/user/database"
	visitDatabase "article-api/models/visit/database"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dsn := username + ":" + password + "@tcp(" + url + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&userDatabase.User{})
	DB.AutoMigrate(&profileDatabase.Profile{})
	DB.AutoMigrate(&articleDatabase.Article{})
	DB.AutoMigrate(&commentDatabase.Comment{})
	DB.AutoMigrate(&likeDatabase.Like{})
	DB.AutoMigrate(&visitDatabase.Visit{})
}

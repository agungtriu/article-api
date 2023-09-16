package database

import (
	articleDatabase "article-api/models/article/database"
	commentDatabase "article-api/models/comment/database"
	likeDatabase "article-api/models/like/database"
	profileDatabase "article-api/models/profile/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Profile  profileDatabase.Profile   `gorm:"foreignKey:UserId;references:ID"`
	Articles []articleDatabase.Article `gorm:"foreignKey:UserId;references:ID"`
	Likes    []likeDatabase.Like       `gorm:"foreignKey:UserId;references:ID"`
	Comments []commentDatabase.Comment `gorm:"foreignKey:UserId;references:ID"`
}

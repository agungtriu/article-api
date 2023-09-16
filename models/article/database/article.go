package database

import (
	commentDatabase "article-api/models/comment/database"
	likeDatabase "article-api/models/like/database"
	visitDatabase "article-api/models/visit/database"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string
	Text     string
	UserId   int
	Likes    []likeDatabase.Like       `gorm:"foreignKey:ArticleId;references:ID"`
	Comments []commentDatabase.Comment `gorm:"foreignKey:ArticleId;references:ID"`
	Visits   []visitDatabase.Visit     `gorm:"foreignKey:ArticleId;references:ID"`
}

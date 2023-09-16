package database

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text      string
	UserId    int
	ArticleId int
}

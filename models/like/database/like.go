package database

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserId    int
	ArticleId int
}

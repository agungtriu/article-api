package database

import "gorm.io/gorm"

type Visit struct {
	gorm.Model
	ArticleId int
}

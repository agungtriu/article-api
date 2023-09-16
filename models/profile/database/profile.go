package database

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name   string
	Bio    string
	UserId int
}

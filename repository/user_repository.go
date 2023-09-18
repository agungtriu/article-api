package repository

import (
	"article-api/models/user/database"
)

type UserRepository interface {
	RegisterUser(user database.User) (database.User, error)
	GetUser(id int) (database.User, error)
	ChangeUsername(userId int, username string) (database.User, error)
	ChangeEmail(userId int, email string) (database.User, error)
	ChangePassword(userId int, password string) (database.User, error)
	VerifyUsername(username string) (database.User, error)
	VerifyEmail(email string) (database.User, error)
}

func (r *repository) RegisterUser(user database.User) (database.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
func (r *repository) GetUser(id int) (database.User, error) {
	var user database.User
	err := r.db.Model(&database.User{}).Preload("Profile").First(&user, "id = ?", id).Error
	return user, err
}
func (r *repository) ChangeUsername(userId int, username string) (database.User, error) {
	var user database.User
	err := r.db.Model(&user).Where("id = ?", userId).Update("username", username).Error
	return user, err
}
func (r *repository) ChangeEmail(userId int, email string) (database.User, error) {
	var user database.User
	err := r.db.Model(&user).Where("id = ?", userId).Update("email", email).Error
	return user, err
}
func (r *repository) ChangePassword(userId int, password string) (database.User, error) {
	var user database.User
	err := r.db.Model(&user).Where("id = ?", userId).Update("password", password).Error
	return user, err
}
func (r *repository) VerifyUsername(username string) (database.User, error) {
	var user database.User
	err := r.db.First(&user, "username = ?", username).Error
	return user, err
}

func (r *repository) VerifyEmail(email string) (database.User, error) {
	var user database.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

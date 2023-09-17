package repository

import (
	"article-api/helper"
	"article-api/models/user/database"
	"article-api/models/user/request"
	"strings"
)

type UserRepository interface {
	RegisterUser(requestRegister request.Register) (database.User, error)
	GetUser(id int) (database.User, error)
	ChangeUsername(userId int, requestChangeUsername request.ChangeUsername) (database.User, error)
	ChangeEmail(userId int, requestChangeEmail request.ChangeEmail) (database.User, error)
	ChangePassword(userId, requestChangePassword request.ChangePassword) (database.User, error)
	VerifyUsername(username string) (database.User, error)
	VerifyEmail(email string) (database.User, error)
}

func (r *repository) RegisterUser(requestRegister request.Register) (database.User, error) {
	hashedPassword, _ := helper.HashPassword(requestRegister.Password)

	user := database.User{Username: strings.ToLower(requestRegister.Username), Email: strings.ToLower(requestRegister.Email), Password: hashedPassword}
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
	err := r.db.Model(&user).Where("id = ?", userId).Update("username", strings.ToLower(username)).Error
	return user, err
}
func (r *repository) ChangeEmail(userId int, requestChangeEmail request.ChangeEmail) (database.User, error) {
	var user database.User
	err := r.db.Model(&user).Where("id = ?", userId).Update("email", strings.ToLower(requestChangeEmail.Email)).Error
	return user, err
}
func (r *repository) ChangePassword(userId int, requestChangePassword request.ChangePassword) (database.User, error) {
	var user database.User
	hashedPassword, _ := helper.HashPassword(requestChangePassword.NewPassword)
	err := r.db.Model(&user).Where("id = ?", userId).Update("password", hashedPassword).Error

	return user, err
}
func (r *repository) VerifyUsername(username string) (database.User, error) {
	var user database.User

	err := r.db.First(&user, "username = ?", strings.ToLower(username)).Error

	return user, err
}

func (r *repository) VerifyEmail(email string) (database.User, error) {
	var user database.User

	err := r.db.First(&user, "email = ?", strings.ToLower(email)).Error

	return user, err
}

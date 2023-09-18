package service

import (
	"article-api/helper"
	"article-api/models/user/database"
	"article-api/models/user/request"
	"article-api/repository"
	"strings"
)

type UserService interface {
	RegisterUser(requestRegister request.Register) (database.User, error)
	GetUser(id int) (database.User, error)
	ChangeUsername(userId int, username string) (database.User, error)
	ChangeEmail(userId int, email string) (database.User, error)
	ChangePassword(userId int, password string) (database.User, error)
	VerifyUsername(username string) (database.User, error)
	VerifyEmail(email string) (database.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) RegisterUser(requestRegister request.Register) (database.User, error) {
	hashedPassword, _ := helper.HashPassword(requestRegister.Password)

	user := database.User{Username: strings.ToLower(requestRegister.Username), Email: strings.ToLower(requestRegister.Email), Password: hashedPassword}
	user, err := s.repository.RegisterUser(user)
	return user, err
}

func (s *userService) GetUser(id int) (database.User, error) {
	user, err := s.repository.GetUser(id)
	return user, err
}

func (s *userService) ChangeUsername(userId int, username string) (database.User, error) {
	user, err := s.repository.ChangeUsername(userId, strings.ToLower(username))
	return user, err
}

func (s *userService) ChangeEmail(userId int, email string) (database.User, error) {
	user, err := s.repository.ChangeEmail(userId, strings.ToLower(email))
	return user, err
}
func (s *userService) ChangePassword(userId int, password string) (database.User, error) {
	hashedPassword, _ := helper.HashPassword(password)
	user, err := s.repository.ChangePassword(userId, hashedPassword)
	return user, err
}
func (s *userService) VerifyUsername(username string) (database.User, error) {
	user, err := s.repository.VerifyUsername(strings.ToLower(username))
	return user, err
}
func (s *userService) VerifyEmail(email string) (database.User, error) {
	user, err := s.repository.VerifyEmail(strings.ToLower(email))
	return user, err
}

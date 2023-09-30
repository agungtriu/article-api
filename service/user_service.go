package service

import (
	"article-api/helper"
	"article-api/models/user/database"
	"article-api/models/user/request"
	"article-api/models/user/response"
	"article-api/repository"
	"strings"
)

type UserService interface {
	RegisterUser(requestRegister request.Register, channel chan response.Result)
	GetUser(id int, channel chan response.Result)
	ChangeUsername(userId int, username string, channel chan response.Result)
	ChangeEmail(userId int, email string, channel chan response.Result)
	ChangePassword(userId int, password string, channel chan response.Result)
	VerifyUsername(username string, channel chan response.Result)
	VerifyEmail(email string, channel chan response.Result)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) RegisterUser(requestRegister request.Register, channel chan response.Result) {
	hashedPassword, _ := helper.HashPassword(requestRegister.Password)

	user := database.User{Username: strings.ToLower(requestRegister.Username), Email: strings.ToLower(requestRegister.Email), Password: hashedPassword}
	user, err := s.repository.RegisterUser(user)
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}

func (s *userService) GetUser(id int, channel chan response.Result) {
	user, err := s.repository.GetUser(id)
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}

func (s *userService) ChangeUsername(userId int, username string, channel chan response.Result) {
	user, err := s.repository.ChangeUsername(userId, strings.ToLower(username))
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}

func (s *userService) ChangeEmail(userId int, email string, channel chan response.Result) {
	user, err := s.repository.ChangeEmail(userId, strings.ToLower(email))
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}
func (s *userService) ChangePassword(userId int, password string, channel chan response.Result) {
	hashedPassword, _ := helper.HashPassword(password)
	user, err := s.repository.ChangePassword(userId, hashedPassword)
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}
func (s *userService) VerifyUsername(username string, channel chan response.Result) {
	user, err := s.repository.VerifyUsername(strings.ToLower(username))
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}
func (s *userService) VerifyEmail(email string, channel chan response.Result) {
	user, err := s.repository.VerifyEmail(strings.ToLower(email))
	res := new(response.Result)
	res.User = user
	res.Err = err
	channel <- *res
}

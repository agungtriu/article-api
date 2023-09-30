package service

import (
	"article-api/models/profile/database"
	"article-api/models/profile/request"
	"article-api/models/profile/response"
	"article-api/repository"
)

type ProfileService interface {
	RegisterProfile(userId int, channel chan response.Result)
	ChangeProfile(userId int, requestChangeProfile request.ChangeProfile, channel chan response.Result)
}

type profileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(repository repository.ProfileRepository) *profileService {
	return &profileService{repository}
}

func (s *profileService) RegisterProfile(userId int, channel chan response.Result) {
	profile := database.Profile{Name: "", Bio: "", UserId: userId}
	profile, err := s.repository.RegisterProfile(profile)
	res := new(response.Result)
	res.Profile = profile
	res.Err = err
	channel <- *res
}

func (s *profileService) ChangeProfile(userId int, requestChangeProfile request.ChangeProfile, channel chan response.Result) {
	profile := database.Profile{Name: requestChangeProfile.Name, Bio: requestChangeProfile.Bio}
	profile, err := s.repository.ChangeProfile(userId, profile)
	res := new(response.Result)
	res.Profile = profile
	res.Err = err
	channel <- *res
}

package service

import (
	"article-api/models/profile/database"
	"article-api/models/profile/request"
	"article-api/repository"
)

type ProfileService interface {
	RegisterProfile(userId int) (database.Profile, error)
	ChangeProfile(userId int, requestChangeProfile request.ChangeProfile) (database.Profile, error)
}

type profileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(repository repository.ProfileRepository) *profileService {
	return &profileService{repository}
}

func (s *profileService) RegisterProfile(userId int) (database.Profile, error) {
	profile := database.Profile{Name: "", Bio: "", UserId: userId}
	profile, err := s.repository.RegisterProfile(profile)
	return profile, err
}

func (s *profileService) ChangeProfile(userId int, requestChangeProfile request.ChangeProfile) (database.Profile, error) {
	profile := database.Profile{Name: requestChangeProfile.Name, Bio: requestChangeProfile.Bio}
	profile, err := s.repository.ChangeProfile(userId, profile)
	return profile, err
}

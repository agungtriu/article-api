package repository

import (
	"article-api/models/profile/database"
	"article-api/models/profile/request"
)

type ProfileRepository interface {
	RegisterProfile(userId int) (database.Profile, error)
	ChangeProfile(userId int, requestChangeProfile request.ChangeProfile) (database.Profile, error)
}

func (r *repository) RegisterProfile(userId int) (database.Profile, error) {
	profile := database.Profile{Name: "", Bio: "", UserId: userId}
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) ChangeProfile(userId int, requestChangeProfile request.ChangeProfile) (database.Profile, error) {
	var profile database.Profile
	err := r.db.Model(&profile).Where("user_id = ?", userId).Updates(database.Profile{Name: requestChangeProfile.Name, Bio: requestChangeProfile.Bio}).Error

	return profile, err
}

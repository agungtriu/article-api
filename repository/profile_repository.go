package repository

import (
	"article-api/models/profile/database"
)

type ProfileRepository interface {
	RegisterProfile(profile database.Profile) (database.Profile, error)
	ChangeProfile(userId int, requestProfile database.Profile) (database.Profile, error)
}

func (r *repository) RegisterProfile(profile database.Profile) (database.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) ChangeProfile(userId int, requestProfile database.Profile) (database.Profile, error) {
	var profile database.Profile
	err := r.db.Model(&profile).Where("user_id = ?", userId).Updates(requestProfile).Error

	return profile, err
}

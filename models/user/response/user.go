package response

import (
	"article-api/models/user/database"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

func (user *User) MapUserFromDatabase(databaseUser database.User) {
	user.Username = databaseUser.Username
	user.Email = databaseUser.Email
	user.Name = databaseUser.Profile.Name
	user.Bio = databaseUser.Profile.Bio
}

package response

import (
	"article-api/models/user/database"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

func (userResponse *UserResponse) MapUserFromDatabase(userDatabase database.User) {
	userResponse.Username = userDatabase.Username
	userResponse.Email = userDatabase.Email
	userResponse.Name = userDatabase.Profile.Name
	userResponse.Bio = userDatabase.Profile.Bio
}

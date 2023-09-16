package response

import (
	"article-api/middlewares"
	"article-api/models/user/database"
)

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (loginResponse *LoginResponse) MapLoginFromDatabase(userDatabase database.User) {
	loginResponse.Username = userDatabase.Username
	loginResponse.Email = userDatabase.Email
	loginResponse.Token = middlewares.GenerateJWT(userDatabase)
}

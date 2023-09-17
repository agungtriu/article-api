package response

import (
	"article-api/middlewares"
	"article-api/models/user/database"
)

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (login *Login) MapLoginFromDatabase(databaseUser database.User) {
	login.Username = databaseUser.Username
	login.Email = databaseUser.Email
	login.Token = middlewares.GenerateJWT(databaseUser)
}

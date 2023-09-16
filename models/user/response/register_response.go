package response

import (
	"article-api/models/user/database"
)

type RegisterResponse struct {
	Id int `json:"id"`
}

func (registerResponse *RegisterResponse) MapRegisterFromDatabase(userDatabase database.User) {
	registerResponse.Id = int(userDatabase.ID)
}

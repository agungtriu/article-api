package response

import (
	"article-api/models/user/database"
)

type Register struct {
	Id int `json:"id"`
}

func (register *Register) MapRegisterFromDatabase(databaseUser database.User) {
	register.Id = int(databaseUser.ID)
}

package response

import (
	"article-api/models/comment/database"
)

type AddComment struct {
	Id uint `json:"id"`
}

func (comment *AddComment) MapAddCommentFromDatabase(databaseComment database.Comment) {
	comment.Id = databaseComment.ID
}

package response

import (
	"article-api/models/comment/database"
	"time"
)

type Comment struct {
	Id        uint      `json:"id"`
	Text      string    `json:"text"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (comment *Comment) MapCommentFromDatabase(databaseComment database.Comment) {
	comment.Id = databaseComment.ID
	comment.Text = databaseComment.Text
	comment.UserId = databaseComment.UserId
	comment.CreatedAt = databaseComment.CreatedAt
	comment.UpdatedAt = databaseComment.UpdatedAt
}

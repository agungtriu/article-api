package response

import (
	"article-api/models/comment/database"
	"time"
)

type CommentResponse struct {
	Id        uint      `json:"id"`
	Text      string    `json:"text"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (commentResponse *CommentResponse) MapCommentFromDatabase(commentDatabase database.Comment) {
	commentResponse.Id = commentDatabase.ID
	commentResponse.Text = commentDatabase.Text
	commentResponse.UserId = commentDatabase.UserId
	commentResponse.CreatedAt = commentDatabase.CreatedAt
	commentResponse.UpdatedAt = commentDatabase.UpdatedAt
}

package response

import (
	articleDatabase "article-api/models/article/database"
	"article-api/models/comment/response"
	"time"
)

type ArticleResponse struct {
	Id        uint                       `json:"id"`
	Title     string                     `json:"title"`
	Text      string                     `json:"text"`
	Like      int                        `json:"like"`
	Visit     int                        `json:"visit"`
	UserId    int                        `json:"userId"`
	CreatedAt time.Time                  `json:"createdAt"`
	UpdatedAt time.Time                  `json:"updatedAt"`
	Comments  []response.CommentResponse `json:"comments"`
}

func (articleResponse *ArticleResponse) MapArticleFromDatabase(articleDatabase articleDatabase.Article) {
	articleResponse.Id = articleDatabase.ID
	articleResponse.Title = articleDatabase.Title
	articleResponse.Text = articleDatabase.Text
	articleResponse.Like = len(articleDatabase.Likes)
	articleResponse.Visit = len(articleDatabase.Visits)
	articleResponse.UserId = articleDatabase.UserId
	articleResponse.CreatedAt = articleDatabase.CreatedAt
	articleResponse.UpdatedAt = articleDatabase.UpdatedAt

	var commentResponse response.CommentResponse
	for _, v := range articleDatabase.Comments {
		commentResponse.MapCommentFromDatabase(v)
		articleResponse.Comments = append(articleResponse.Comments, commentResponse)
	}
}

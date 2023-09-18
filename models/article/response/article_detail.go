package response

import (
	"article-api/models/article/database"
	"article-api/models/comment/response"
	"time"
)

type ArticleDetail struct {
	Id        uint               `json:"id"`
	Title     string             `json:"title"`
	Text      string             `json:"text"`
	Like      int                `json:"like"`
	Visit     int                `json:"visit"`
	UserId    int                `json:"userId"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Comments  []response.Comment `json:"comments"`
}

func (article *ArticleDetail) MapArticleDetailFromDatabase(databaseArticle database.Article) {
	article.Id = databaseArticle.ID
	article.Title = databaseArticle.Title
	article.Text = databaseArticle.Text
	article.Like = len(databaseArticle.Likes)
	article.Visit = len(databaseArticle.Visits)
	article.UserId = databaseArticle.UserId
	article.CreatedAt = databaseArticle.CreatedAt
	article.UpdatedAt = databaseArticle.UpdatedAt

	var responseComment response.Comment
	for _, v := range databaseArticle.Comments {
		responseComment.MapCommentFromDatabase(v)
		article.Comments = append(article.Comments, responseComment)
	}
}

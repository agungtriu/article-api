package response

import (
	"article-api/models/article/database"
	"time"
)

type Article struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Like      int       `json:"like"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (article *Article) MapArticleFromDatabase(databaseArticle database.Article) {
	article.Id = databaseArticle.ID
	article.Title = databaseArticle.Title
	article.Text = databaseArticle.Text
	article.Like = len(databaseArticle.Likes)
	article.UserId = databaseArticle.UserId
	article.CreatedAt = databaseArticle.CreatedAt
	article.UpdatedAt = databaseArticle.UpdatedAt
}

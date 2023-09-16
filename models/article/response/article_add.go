package response

import (
	"article-api/models/article/database"
)

type AddArticleResponse struct {
	Id int `json:"id"`
}

func (addArticleResponse *AddArticleResponse) MapAddArticleFromDatabase(articleDatabase database.Article) {
	addArticleResponse.Id = int(articleDatabase.ID)
}

package response

import (
	"article-api/models/article/database"
)

type AddArticle struct {
	Id int `json:"id"`
}

func (addArticle *AddArticle) MapAddArticleFromDatabase(databaseArticle database.Article) {
	addArticle.Id = int(databaseArticle.ID)
}

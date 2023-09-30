package response

import "article-api/models/article/database"

type Results struct {
	Articles []database.Article
	Err      error
}

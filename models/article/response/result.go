package response

import "article-api/models/article/database"

type Result struct {
	Article database.Article
	Err     error
}

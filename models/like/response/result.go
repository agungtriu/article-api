package response

import "article-api/models/like/database"

type Result struct {
	Like database.Like
	Err  error
}

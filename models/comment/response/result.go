package response

import "article-api/models/comment/database"

type Result struct {
	Comment database.Comment
	Err     error
}

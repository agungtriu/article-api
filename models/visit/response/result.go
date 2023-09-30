package response

import "article-api/models/visit/database"

type Result struct {
	Visit database.Visit
	Err   error
}

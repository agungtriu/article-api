package response

import "article-api/models/user/database"

type Result struct {
	User database.User
	Err  error
}

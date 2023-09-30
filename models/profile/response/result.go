package response

import "article-api/models/profile/database"

type Result struct {
	Profile database.Profile
	Err     error
}

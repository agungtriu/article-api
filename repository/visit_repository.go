package repository

import "article-api/models/visit/database"

type VisitRepository interface {
	PostVisit(articleId int) error
}

func (r *repository) PostVisit(articleId int) error {
	visitdatabase := database.Visit{ArticleId: articleId}
	err := r.db.Create(&visitdatabase).Error
	return err
}

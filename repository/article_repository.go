package repository

import (
	"article-api/models/article/database"
)

type ArticleRepository interface {
	GetArticles() ([]database.Article, error)
	SearchArticles(search string) ([]database.Article, error)
	GetArticle(articleId int) (database.Article, error)
	PostArticle(userId int, article database.Article) (database.Article, error)
	PutArticle(userId int, articleId int, article database.Article) (database.Article, error)
	DeleteArticle(userId int, articleId int) error
	VerifyArticle(articleId int) (database.Article, error)
}

func (r *repository) GetArticles() ([]database.Article, error) {
	var articles []database.Article
	err := r.db.Preload("Likes").Find(&articles).Error

	return articles, err
}

func (r *repository) SearchArticles(search string) ([]database.Article, error) {
	var articles []database.Article
	err := r.db.Preload("Likes").Where("title LIKE ? OR text LIKE ?", "%"+search+"%", "%"+search+"%").Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticle(articleId int) (database.Article, error) {
	var article database.Article
	err := r.db.Preload("Likes").Preload("Comments").Preload("Visits").Find(&article, "articles.id = ?", articleId).Error

	return article, err
}
func (r *repository) PostArticle(userId int, article database.Article) (database.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}
func (r *repository) PutArticle(userId int, articleId int, article database.Article) (database.Article, error) {
	err := r.db.Model(&article).Where("id = ? AND user_id = ?", articleId, userId).Updates(article).Error

	return article, err
}
func (r *repository) DeleteArticle(userId int, articleId int) error {
	var article database.Article
	err := r.db.Where("id = ? AND user_id = ?", articleId, userId).Delete(&article).Error
	return err
}
func (r *repository) VerifyArticle(articleId int) (database.Article, error) {
	var article database.Article
	err := r.db.First(&article, "id = ?", articleId).Error

	return article, err
}

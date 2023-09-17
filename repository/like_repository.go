package repository

import (
	"article-api/models/like/database"
)

type LikeRepository interface {
	PostLike(userId int, articleId int) (database.Like, error)

	DeleteLike(userId int, articleId int) error

	VerifyLike(articleId int, userId int) (database.Like, error)
}

func (r *repository) PostLike(userId int, articleId int) (database.Like, error) {
	like := database.Like{UserId: userId, ArticleId: articleId}
	err := r.db.Create(&like).Error

	return like, err
}

func (r *repository) DeleteLike(userId int, articleId int) error {
	var like database.Like
	err := r.db.Where("user_id = ? AND article_id = ?", userId, articleId).Delete(&like).Error

	return err
}

func (r *repository) VerifyLike(articleId int, userId int) (database.Like, error) {
	var like database.Like
	err := r.db.Find(&like, "article_id = ? AND user_id = ?", articleId, userId).Error

	return like, err
}

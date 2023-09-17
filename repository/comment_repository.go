package repository

import (
	"article-api/models/comment/database"
	"article-api/models/comment/request"
)

type CommentRepository interface {
	PostComment(userId int, articleId int, requestComment request.Comment) (database.Comment, error)
	PutComment(commentId int, userId int, articleId int, requestComment request.Comment) (database.Comment, error)
	DeleteComment(commentId int, userId int, articleId int, requestComment request.Comment) error
	VerifyComment(commentId int) (database.Comment, error)
}

func (r *repository) PostComment(userId int, articleId int, requestComment request.Comment) (database.Comment, error) {
	comment := database.Comment{Text: requestComment.Text, UserId: userId, ArticleId: articleId}

	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *repository) PutComment(commentId int, userId int, articleId int, requestComment request.Comment) (database.Comment, error) {
	var comment database.Comment
	err := r.db.Model(&comment).Where("id = ? AND user_id = ? AND article_id = ?", commentId, userId, articleId).Update("text", requestComment.Text).Error

	return comment, err
}

func (r *repository) DeleteComment(commentId int, userId int, articleId int) error {
	var comment database.Comment
	err := r.db.Where("id = ? AND article_id = ? AND user_id = ?", commentId, articleId, userId).Delete(&comment).Error

	return err
}

func (r *repository) VerifyComment(commentId int) (database.Comment, error) {
	var comment database.Comment
	err := r.db.First(&comment, "id = ?", commentId).Error
	return comment, err
}

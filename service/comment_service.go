package service

import (
	"article-api/models/comment/database"
	"article-api/models/comment/request"
	"article-api/repository"
)

type CommentService interface {
	PostComment(userId int, articleId int, requestComment request.Comment) (database.Comment, error)
	PutComment(commentId int, userId int, articleId int, text string) (database.Comment, error)
	DeleteComment(commentId int, userId int, articleId int) error
	VerifyComment(commentId int) (database.Comment, error)
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository}
}

func (s *commentService) PostComment(userId int, articleId int, requestComment request.Comment) (database.Comment, error) {
	comment := database.Comment{Text: requestComment.Text, UserId: userId, ArticleId: articleId}
	comment, err := s.repository.PostComment(comment)
	return comment, err
}

func (s *commentService) PutComment(commentId int, userId int, articleId int, text string) (database.Comment, error) {
	comment, err := s.repository.PutComment(commentId, userId, articleId, text)
	return comment, err
}

func (s *commentService) DeleteComment(commentId int, userId int, articleId int) error {
	err := s.repository.DeleteComment(commentId, userId, articleId)
	return err
}

func (s *commentService) VerifyComment(commentId int) (database.Comment, error) {
	comment, err := s.repository.VerifyComment(commentId)
	return comment, err
}

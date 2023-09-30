package service

import (
	"article-api/models/comment/database"
	"article-api/models/comment/request"
	"article-api/models/comment/response"
	"article-api/repository"
)

type CommentService interface {
	PostComment(userId int, articleId int, requestComment request.Comment, channel chan response.Result)
	PutComment(commentId int, userId int, articleId int, text string, channel chan response.Result)
	DeleteComment(commentId int, userId int, articleId int, channel chan response.Result)
	VerifyComment(commentId int, channel chan response.Result)
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository}
}

func (s *commentService) PostComment(userId int, articleId int, requestComment request.Comment, channel chan response.Result) {
	comment := database.Comment{Text: requestComment.Text, UserId: userId, ArticleId: articleId}
	comment, err := s.repository.PostComment(comment)
	res := new(response.Result)
	res.Comment = comment
	res.Err = err
	channel <- *res
}

func (s *commentService) PutComment(commentId int, userId int, articleId int, text string, channel chan response.Result) {
	comment, err := s.repository.PutComment(commentId, userId, articleId, text)
	res := new(response.Result)
	res.Comment = comment
	res.Err = err
	channel <- *res
}

func (s *commentService) DeleteComment(commentId int, userId int, articleId int, channel chan response.Result) {
	err := s.repository.DeleteComment(commentId, userId, articleId)
	res := new(response.Result)
	res.Err = err
	channel <- *res
}

func (s *commentService) VerifyComment(commentId int, channel chan response.Result) {
	comment, err := s.repository.VerifyComment(commentId)
	res := new(response.Result)
	res.Comment = comment
	res.Err = err
	channel <- *res
}

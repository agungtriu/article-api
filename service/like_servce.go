package service

import (
	"article-api/models/like/response"
	"article-api/repository"
)

type LikeService interface {
	PostLike(userId int, articleId int, channel chan response.Result)
	DeleteLike(userId int, articleId int, channel chan response.Result)
	VerifyLike(articleId int, userId int, channel chan response.Result)
}

type likeService struct {
	repository repository.LikeRepository
}

func NewLikeService(repository repository.LikeRepository) *likeService {
	return &likeService{repository}
}

func (s *likeService) PostLike(userId int, articleId int, channel chan response.Result) {
	like, err := s.repository.PostLike(userId, articleId)
	res := new(response.Result)
	res.Like = like
	res.Err = err
	channel <- *res
}

func (s *likeService) DeleteLike(userId int, articleId int, channel chan response.Result) {
	err := s.repository.DeleteLike(userId, articleId)
	res := new(response.Result)
	res.Err = err
	channel <- *res
}

func (s *likeService) VerifyLike(articleId int, userId int, channel chan response.Result) {
	like, err := s.repository.VerifyLike(articleId, userId)
	res := new(response.Result)
	res.Like = like
	res.Err = err
	channel <- *res
}

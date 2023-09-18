package service

import (
	"article-api/models/like/database"
	"article-api/repository"
)

type LikeService interface {
	PostLike(userId int, articleId int) (database.Like, error)
	DeleteLike(userId int, articleId int) error
	VerifyLike(articleId int, userId int) (database.Like, error)
}

type likeService struct {
	repository repository.LikeRepository
}

func NewLikeService(repository repository.LikeRepository) *likeService {
	return &likeService{repository}
}

func (s *likeService) PostLike(userId int, articleId int) (database.Like, error) {
	like, err := s.repository.PostLike(userId, articleId)
	return like, err
}

func (s *likeService) DeleteLike(userId int, articleId int) error {
	err := s.repository.DeleteLike(userId, articleId)
	return err
}

func (s *likeService) VerifyLike(articleId int, userId int) (database.Like, error) {
	like, err := s.repository.VerifyLike(articleId, userId)
	return like, err
}

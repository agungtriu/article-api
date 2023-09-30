package service

import (
	"article-api/repository"
)

type VisitService interface {
	PostVisit(articleId int)
}

type visitService struct {
	repository repository.VisitRepository
}

func NewVisitService(repository repository.VisitRepository) *visitService {
	return &visitService{repository}
}
func (s *visitService) PostVisit(articleId int) {
	s.repository.PostVisit(articleId)
}

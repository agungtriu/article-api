package service

import "article-api/repository"

type VisitService interface {
	PostVisit(articleId int) error
}

type visitService struct {
	repository repository.VisitRepository
}

func NewVisitService(repository repository.VisitRepository) *visitService {
	return &visitService{repository}
}
func (s *visitService) PostVisit(articleId int) error {
	err := s.repository.PostVisit(articleId)
	return err
}

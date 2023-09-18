package service

import (
	"article-api/models/article/database"
	"article-api/models/article/request"
	"article-api/repository"
)

type ArticleService interface {
	GetArticles() ([]database.Article, error)
	SearchArticles(search string) ([]database.Article, error)
	GetArticle(articleId int) (database.Article, error)
	PostArticle(userId int, requestArticle request.Article) (database.Article, error)
	PutArticle(userId int, articleId int, requestArticle request.Article) (database.Article, error)
	DeleteArticle(userId int, articleId int) error
	VerifyArticle(articleId int) (database.Article, error)
}

type articleService struct {
	repository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) *articleService {
	return &articleService{repository}
}

func (s *articleService) GetArticles() ([]database.Article, error) {
	articles, err := s.repository.GetArticles()
	return articles, err
}

func (s *articleService) SearchArticles(search string) ([]database.Article, error) {
	articles, err := s.repository.SearchArticles(search)
	return articles, err
}

func (s *articleService) GetArticle(articleId int) (database.Article, error) {
	article, err := s.repository.GetArticle(articleId)
	return article, err
}

func (s *articleService) PostArticle(userId int, requestArticle request.Article) (database.Article, error) {
	article := database.Article{Title: requestArticle.Title, Text: requestArticle.Text, UserId: userId}
	article, err := s.repository.PostArticle(userId, article)
	return article, err
}

func (s *articleService) PutArticle(userId int, articleId int, requestArticle request.Article) (database.Article, error) {
	article := database.Article{Title: requestArticle.Title, Text: requestArticle.Text}
	article, err := s.repository.PutArticle(userId, articleId, article)
	return article, err
}

func (s *articleService) DeleteArticle(userId int, articleId int) error {
	err := s.repository.DeleteArticle(userId, articleId)
	return err
}

func (s *articleService) VerifyArticle(articleId int) (database.Article, error) {
	article, err := s.repository.VerifyArticle(articleId)
	return article, err
}

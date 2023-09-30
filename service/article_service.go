package service

import (
	"article-api/models/article/database"
	"article-api/models/article/request"
	"article-api/models/article/response"
	"article-api/repository"
)

type ArticleService interface {
	GetArticles(channel chan response.Results)
	SearchArticles(search string, channel chan response.Results)
	GetArticle(articleId int, channel chan response.Result)
	PostArticle(userId int, requestArticle request.Article, channel chan response.Result)
	PutArticle(userId int, articleId int, requestArticle request.Article, channel chan response.Result)
	DeleteArticle(userId int, articleId int, channel chan response.Result)
	VerifyArticle(articleId int, channel chan response.Result)
}

type articleService struct {
	repository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) *articleService {
	return &articleService{repository}
}

func (s *articleService) GetArticles(channel chan response.Results) {
	articles, err := s.repository.GetArticles()

	res := new(response.Results)
	res.Articles = articles
	res.Err = err
	channel <- *res
}

func (s *articleService) SearchArticles(search string, channel chan response.Results) {
	articles, err := s.repository.SearchArticles(search)

	res := new(response.Results)
	res.Articles = articles
	res.Err = err
	channel <- *res
}

func (s *articleService) GetArticle(articleId int, channel chan response.Result) {
	article, err := s.repository.GetArticle(articleId)

	res := new(response.Result)
	res.Article = article
	res.Err = err
	channel <- *res
}

func (s *articleService) PostArticle(userId int, requestArticle request.Article, channel chan response.Result) {
	article := database.Article{Title: requestArticle.Title, Text: requestArticle.Text, UserId: userId}
	article, err := s.repository.PostArticle(userId, article)

	res := new(response.Result)
	res.Article = article
	res.Err = err
	channel <- *res
}

func (s *articleService) PutArticle(userId int, articleId int, requestArticle request.Article, channel chan response.Result) {
	article := database.Article{Title: requestArticle.Title, Text: requestArticle.Text}
	article, err := s.repository.PutArticle(userId, articleId, article)

	res := new(response.Result)
	res.Article = article
	res.Err = err
	channel <- *res
}

func (s *articleService) DeleteArticle(userId int, articleId int, channel chan response.Result) {
	err := s.repository.DeleteArticle(userId, articleId)
	res := new(response.Result)
	res.Err = err
	channel <- *res
}

func (s *articleService) VerifyArticle(articleId int, channel chan response.Result) {
	article, err := s.repository.VerifyArticle(articleId)
	res := new(response.Result)
	res.Article = article
	res.Err = err
	channel <- *res
}

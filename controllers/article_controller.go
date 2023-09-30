package controllers

import (
	"article-api/middlewares"
	"article-api/models/article/request"
	"article-api/models/article/response"
	"article-api/models/base"
	"article-api/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type articleController struct {
	articleService service.ArticleService
	visitService   service.VisitService
}

func NewArticleController(articleService service.ArticleService, visitService service.VisitService) *articleController {
	return &articleController{articleService, visitService}
}

func (controller *articleController) GetArticlesController(c echo.Context) error {
	search := c.QueryParam("search")

	channel := make(chan response.Results)

	if search != "" {
		go controller.articleService.SearchArticles(search, channel)
	} else {
		go controller.articleService.GetArticles(channel)
	}

	articles := <-channel

	if articles.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: true,
			Error:  articles.Err.Error(),
		})
	}

	var responseArticle response.Article
	var responseArticles []response.Article
	for _, v := range articles.Articles {
		responseArticle.MapArticleFromDatabase(v)
		responseArticles = append(responseArticles, responseArticle)
	}

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data articles",
		Data:    responseArticles,
	})
}

func (controller *articleController) GetArticleController(c echo.Context) error {
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	go controller.visitService.PostVisit(articleId)

	channel := make(chan response.Result)
	go controller.articleService.GetArticle(articleId, channel)

	article := <-channel

	if article.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  article.Err.Error(),
		})
	}

	if article.Article.ID == 0 {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var responseArticle response.ArticleDetail
	responseArticle.MapArticleDetailFromDatabase(article.Article)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data article",
		Data:    responseArticle,
	})
}

func (controller *articleController) AddArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestArticle request.Article
	c.Bind(&requestArticle)

	channel := make(chan response.Result)

	go controller.articleService.PostArticle(userId, requestArticle, channel)

	article := <-channel

	if article.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  article.Err.Error(),
		})
	}

	var reponseArticle response.AddArticle
	reponseArticle.MapAddArticleFromDatabase(article.Article)

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success created article",
		Data:    reponseArticle,
	})
}

func (controller *articleController) UpdateArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	channel := make(chan response.Result)
	go controller.articleService.VerifyArticle(articleId, channel)

	article := <-channel
	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.Article.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	var requestArticle request.Article
	c.Bind(&requestArticle)

	go controller.articleService.PutArticle(userId, articleId, requestArticle, channel)

	article = <-channel

	if article.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  article.Err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success updated article",
	})
}

func (controller *articleController) DeleteArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	channel := make(chan response.Result)

	go controller.articleService.VerifyArticle(articleId, channel)

	article := <-channel

	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.Article.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	go controller.articleService.DeleteArticle(userId, articleId, channel)
	article = <-channel
	if article.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  article.Err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted article",
	})
}

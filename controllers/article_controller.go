package controllers

import (
	"article-api/middlewares"
	"article-api/models/article/database"
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

	var articles []database.Article
	var err error
	if search != "" {
		articles, err = controller.articleService.SearchArticles(search)
	} else {
		articles, err = controller.articleService.GetArticles()
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: true,
			Error:  err.Error(),
		})
	}

	var responseArticle response.Article
	var responseArticles []response.Article
	for _, v := range articles {
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

	controller.visitService.PostVisit(articleId)

	article, err := controller.articleService.GetArticle(articleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	if article.ID == 0 {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var responseArticle response.ArticleDetail
	responseArticle.MapArticleDetailFromDatabase(article)

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

	article, err := controller.articleService.PostArticle(userId, requestArticle)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	var reponseArticle response.AddArticle
	reponseArticle.MapAddArticleFromDatabase(article)

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

	article, err := controller.articleService.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	var requestArticle request.Article
	c.Bind(&requestArticle)

	article, err = controller.articleService.PutArticle(userId, articleId, requestArticle)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
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

	article, err := controller.articleService.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	err = controller.articleService.DeleteArticle(userId, articleId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted article",
	})
}

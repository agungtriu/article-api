package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/article/database"
	"article-api/models/article/request"
	"article-api/models/article/response"
	"article-api/models/base"
	"article-api/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetArticlesController(c echo.Context) error {
	search := c.QueryParam("search")

	var articles []database.Article
	var err error
	repository := repository.NewRepository(configs.DB)
	if search != "" {
		articles, err = repository.SearchArticles(search)
	} else {
		articles, err = repository.GetArticles()
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

func GetArticleController(c echo.Context) error {
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	repository := repository.NewRepository(configs.DB)
	repository.PostVisit(articleId)

	article, err := repository.GetArticle(articleId)
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

func AddArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestArticle request.Article
	c.Bind(&requestArticle)

	repository := repository.NewRepository(configs.DB)
	article, err := repository.PostArticle(userId, requestArticle)

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

func UpdateArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	repository := repository.NewRepository(configs.DB)
	article, err := repository.VerifyArticle(articleId)

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

	article, err = repository.PutArticle(userId, articleId, requestArticle)

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

func DeleteArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	repository := repository.NewRepository(configs.DB)
	article, err := repository.VerifyArticle(articleId)

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

	err = repository.DeleteArticle(userId, articleId)

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

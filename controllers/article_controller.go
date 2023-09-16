package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	articledatabase "article-api/models/article/database"
	"article-api/models/article/request"
	"article-api/models/article/response"
	"article-api/models/base"
	visitdatabase "article-api/models/visit/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetArticlesController(c echo.Context) error {
	search := c.QueryParam("search")

	var articles []articledatabase.Article
	var result *gorm.DB

	if search != "" {
		result = configs.DB.Preload("Likes").Preload("Comments").Preload("Visits").Where("title LIKE ? OR text LIKE ?", "%"+search+"%", "%"+search+"%").Find(&articles)
	} else {
		result = configs.DB.Preload("Likes").Preload("Comments").Preload("Visits").Find(&articles)
	}

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: true,
			Error:  result.Error,
		})
	}

	var articleResponse response.ArticleResponse
	var resultResponse []response.ArticleResponse
	for _, v := range articles {
		articleResponse.MapArticleFromDatabase(v)
		resultResponse = append(resultResponse, articleResponse)
	}

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data articles",
		Data:    resultResponse,
	})
}

func GetArticleController(c echo.Context) error {
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	visitdatabase := visitdatabase.Visit{ArticleId: articleId}
	configs.DB.Create(&visitdatabase)

	var article articledatabase.Article
	result := configs.DB.Preload("Likes").Preload("Comments").Preload("Visits").Find(&article, "articles.id = ?", articleId)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	// Maping to response
	var articleResponse response.ArticleResponse
	articleResponse.MapArticleFromDatabase(article)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data article",
		Data:    articleResponse,
	})
}

func AddArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var articleRequest request.Article
	c.Bind(&articleRequest)
	article := articledatabase.Article{Title: articleRequest.Title, Text: articleRequest.Text, UserId: userId}

	result := configs.DB.Create(&article)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	var articleResponse response.AddArticleResponse
	articleResponse.MapAddArticleFromDatabase(article)

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success created article",
		Data:    articleResponse,
	})
}

func UpdateArticleController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId := c.Param("articleId")

	var article articledatabase.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.UserId {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	var articleRequest request.Article
	c.Bind(&articleRequest)

	result := configs.DB.Model(&article).Where("id = ? AND user_id = ?", articleId, userId).Updates(articledatabase.Article{Title: articleRequest.Title, Text: articleRequest.Text})

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
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
	articleId := c.Param("articleId")

	var article articledatabase.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	if userId != article.UserId {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	result := configs.DB.Where("id = ? AND user_id = ?", articleId, userId).Delete(&article)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted article",
	})
}

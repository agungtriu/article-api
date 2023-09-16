package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	articledatabase "article-api/models/article/database"
	"article-api/models/base"
	likedatabase "article-api/models/like/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddLikeController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	var article articledatabase.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var likeDatabase likedatabase.Like
	configs.DB.Find(&likeDatabase, "article_id = ? AND user_id = ?", articleId, userId)

	if likeDatabase.ID != 0 {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Cannot double like article",
		})
	}

	like := likedatabase.Like{UserId: userId, ArticleId: articleId}
	result := configs.DB.Create(&like)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success like article",
	})
}

func DeleteLikeController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	var article articledatabase.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var likeDatabase likedatabase.Like
	configs.DB.Find(&likeDatabase, "article_id = ? AND user_id = ?", articleId, userId)

	if likeDatabase.ID == 0 {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Like not found",
		})
	}

	var like likedatabase.Like
	result := configs.DB.Where("user_id = ? AND article_id = ?", userId, articleId).Delete(&like)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success unlike article",
	})
}

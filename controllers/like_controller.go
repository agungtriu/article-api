package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/base"
	"article-api/repository"
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

	repository := repository.NewRepository(configs.DB)
	_, err := repository.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	like, _ := repository.VerifyLike(articleId, userId)

	if like.ID != 0 {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Cannot double like article",
		})
	}

	like, err = repository.PostLike(userId, articleId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
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

	repository := repository.NewRepository(configs.DB)
	_, err := repository.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	like, _ := repository.VerifyLike(articleId, userId)

	if like.ID == 0 {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Like not found",
		})
	}

	err = repository.DeleteLike(userId, articleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success unlike article",
	})
}

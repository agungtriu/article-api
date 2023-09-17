package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/base"
	"article-api/models/comment/request"
	"article-api/models/comment/response"
	"article-api/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	repository := repository.NewRepository(configs.DB)
	_, err := repository.GetArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var commentRequest request.Comment
	c.Bind(&commentRequest)

	comment, err := repository.PostComment(userId, articleId, commentRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	var responseComment response.AddComment
	responseComment.MapAddCommentFromDatabase(comment)

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success created comment",
		Data:    responseComment,
	})
}

func UpdateCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	var requestComment request.Comment
	c.Bind(&requestComment)

	repository := repository.NewRepository(configs.DB)
	_, err := repository.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	comment, err := repository.VerifyComment(commentId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Comment not found",
		})
	}

	if userId != comment.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	comment, err = repository.PutComment(commentId, userId, articleId, requestComment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success updated comment",
	})
}

func DeleteCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	repository := repository.NewRepository(configs.DB)

	_, err := repository.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	comment, err := repository.VerifyComment(commentId)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Comment not found",
		})
	}

	if userId != comment.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	err = repository.DeleteComment(commentId, userId, articleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted comment",
	})
}

package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/article/database"
	"article-api/models/base"
	commentdatabase "article-api/models/comment/database"
	"article-api/models/comment/request"
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

	var article database.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var commentRequest request.Comment
	c.Bind(&commentRequest)

	comment := commentdatabase.Comment{Text: commentRequest.Text, UserId: userId, ArticleId: articleId}

	result := configs.DB.Create(&comment)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success created comment",
		Data:    map[string]uint{"id": comment.ID},
	})
}

func UpdateCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId := c.Param("articleId")
	commentId := c.Param("commentId")

	var commentRequest request.Comment
	c.Bind(&commentRequest)

	var article database.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var comment commentdatabase.Comment
	verifyComment := configs.DB.First(&comment, "id = ?", commentId)

	if verifyComment.Error != nil {
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

	result := configs.DB.Model(&comment).Where("id = ? AND user_id = ? AND article_id = ?", commentId, userId, articleId).Update("text", commentRequest.Text)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
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
	articleId := c.Param("articleId")
	commentId := c.Param("commentId")

	var article database.Article
	verifyArticle := configs.DB.First(&article, "id = ?", articleId)

	if verifyArticle.Error != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var comment commentdatabase.Comment
	verifyComment := configs.DB.First(&comment, "id = ?", commentId)

	if verifyComment.Error != nil {
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

	result := configs.DB.Where("id = ? AND article_id = ? AND user_id = ?", commentId, articleId, userId).Delete(&comment)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted comment",
	})
}

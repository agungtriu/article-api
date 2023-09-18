package controllers

import (
	"article-api/middlewares"
	"article-api/models/base"
	"article-api/models/comment/request"
	"article-api/models/comment/response"
	"article-api/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type commentController struct {
	commentService service.CommentService
	articleService service.ArticleService
}

func NewCommentController(commentService service.CommentService, articleService service.ArticleService) *commentController {
	return &commentController{commentService, articleService}
}
func (controller *commentController) AddCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	_, err := controller.articleService.GetArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var commentRequest request.Comment
	c.Bind(&commentRequest)

	comment, err := controller.commentService.PostComment(userId, articleId, commentRequest)

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

func (controller *commentController) UpdateCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	var requestComment request.Comment
	c.Bind(&requestComment)

	_, err := controller.articleService.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	comment, err := controller.commentService.VerifyComment(commentId)

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

	comment, err = controller.commentService.PutComment(commentId, userId, articleId, requestComment.Text)
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

func (controller *commentController) DeleteCommentController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	_, err := controller.articleService.VerifyArticle(articleId)

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	comment, err := controller.commentService.VerifyComment(commentId)
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

	err = controller.commentService.DeleteComment(commentId, userId, articleId)
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

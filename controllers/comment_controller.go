package controllers

import (
	"article-api/middlewares"
	"article-api/models/article/response"
	"article-api/models/base"
	"article-api/models/comment/request"
	commentResponse "article-api/models/comment/response"
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

	channelArticle := make(chan response.Result)

	go controller.articleService.GetArticle(articleId, channelArticle)

	article := <-channelArticle

	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	var commentRequest request.Comment
	c.Bind(&commentRequest)

	channelComment := make(chan commentResponse.Result)
	go controller.commentService.PostComment(userId, articleId, commentRequest, channelComment)

	comment := <-channelComment

	if comment.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  comment.Err.Error(),
		})
	}

	var responseComment commentResponse.AddComment
	responseComment.MapAddCommentFromDatabase(comment.Comment)

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

	channel := make(chan response.Result)

	go controller.articleService.VerifyArticle(articleId, channel)

	article := <-channel
	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	channelComment := make(chan commentResponse.Result)
	go controller.commentService.VerifyComment(commentId, channelComment)

	comment := <-channelComment
	if comment.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Comment not found",
		})
	}

	if userId != comment.Comment.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	go controller.commentService.PutComment(commentId, userId, articleId, requestComment.Text, channelComment)
	comment = <-channelComment
	if comment.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  comment.Err.Error(),
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

	channel := make(chan response.Result)
	go controller.articleService.VerifyArticle(articleId, channel)

	article := <-channel

	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	channelComment := make(chan commentResponse.Result)
	go controller.commentService.VerifyComment(commentId, channelComment)
	comment := <-channelComment
	if comment.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Comment not found",
		})
	}

	if userId != comment.Comment.UserId {
		return c.JSON(http.StatusForbidden, base.ErrorResponse{
			Status: false,
			Error:  "Unauthorized Access",
		})
	}

	go controller.commentService.DeleteComment(commentId, userId, articleId, channelComment)
	comment = <-channelComment
	if comment.Err != nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  comment.Err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted comment",
	})
}

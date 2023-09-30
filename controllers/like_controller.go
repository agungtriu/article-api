package controllers

import (
	"article-api/middlewares"
	articleResponse "article-api/models/article/response"
	"article-api/models/base"
	likeResponse "article-api/models/like/response"
	"article-api/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type likeController struct {
	likeService    service.LikeService
	articleService service.ArticleService
}

func NewLikeController(likeService service.LikeService, articleService service.ArticleService) *likeController {
	return &likeController{likeService, articleService}
}

func (controller *likeController) AddLikeController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	channelArticle := make(chan articleResponse.Result)
	go controller.articleService.VerifyArticle(articleId, channelArticle)

	article := <-channelArticle

	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}
	channelLike := make(chan likeResponse.Result)
	go controller.likeService.VerifyLike(articleId, userId, channelLike)
	like := <-channelLike

	if like.Like.ID != 0 {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Cannot double like article",
		})
	}

	go controller.likeService.PostLike(userId, articleId, channelLike)
	like = <-channelLike

	if like.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  like.Err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success like article",
	})
}

func (controller *likeController) DeleteLikeController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])

	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	articleId, _ := strconv.Atoi(c.Param("articleId"))

	channelArticle := make(chan articleResponse.Result)
	go controller.articleService.VerifyArticle(articleId, channelArticle)

	article := <-channelArticle

	if article.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Article not found",
		})
	}

	channelLike := make(chan likeResponse.Result)
	go controller.likeService.VerifyLike(articleId, userId, channelLike)
	like := <-channelLike

	if like.Like.ID == 0 {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "Like not found",
		})
	}

	go controller.likeService.DeleteLike(userId, articleId, channelLike)
	like = <-channelLike
	if like.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  like.Err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success unlike article",
	})
}

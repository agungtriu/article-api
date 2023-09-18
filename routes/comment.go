package routes

import (
	"article-api/configs"
	"article-api/controllers"
	"article-api/repository"
	"article-api/service"

	"github.com/labstack/echo/v4"
)

func CommentRoute(eAuthArticle *echo.Group) {
	repository := repository.NewRepository(configs.DB)
	commentService := service.NewCommentService(repository)
	articleService := service.NewArticleService(repository)
	commentController := controllers.NewCommentController(commentService, articleService)

	eAuthComments := eAuthArticle.Group("/comments")
	eAuthComments.POST("", commentController.AddCommentController)
	eAuthComments.PUT("/:commentId", commentController.UpdateCommentController)
	eAuthComments.DELETE("/:commentId", commentController.DeleteCommentController)
}

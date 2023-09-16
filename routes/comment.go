package routes

import (
	"article-api/controllers"

	"github.com/labstack/echo/v4"
)

func CommentRoute(eAuthArticle *echo.Group) {
	eAuthComments := eAuthArticle.Group("/comments")
	eAuthComments.POST("", controllers.AddCommentController)
	eAuthComments.PUT("/:commentId", controllers.UpdateCommentController)
	eAuthComments.DELETE("/:commentId", controllers.DeleteCommentController)
}

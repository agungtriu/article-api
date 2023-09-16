package routes

import (
	"article-api/controllers"

	"github.com/labstack/echo/v4"
)

func LikeRoute(eAuthArticle *echo.Group) {
	eAuthLikes := eAuthArticle.Group("/likes")
	eAuthLikes.POST("", controllers.AddLikeController)
	eAuthLikes.DELETE("", controllers.DeleteLikeController)
}

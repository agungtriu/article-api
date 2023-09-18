package routes

import (
	"article-api/configs"
	"article-api/controllers"
	"article-api/repository"
	"article-api/service"

	"github.com/labstack/echo/v4"
)

func LikeRoute(eAuthArticle *echo.Group) {
	repository := repository.NewRepository(configs.DB)
	likeService := service.NewLikeService(repository)
	articleService := service.NewArticleService(repository)
	likeController := controllers.NewLikeController(likeService, articleService)

	eAuthLikes := eAuthArticle.Group("/likes")
	eAuthLikes.POST("", likeController.AddLikeController)
	eAuthLikes.DELETE("", likeController.DeleteLikeController)
}

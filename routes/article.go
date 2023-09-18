package routes

import (
	"article-api/configs"
	"article-api/controllers"
	"article-api/repository"
	"article-api/service"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ArticleRoute(e *echo.Echo) {
	repository := repository.NewRepository(configs.DB)
	articleService := service.NewArticleService(repository)
	visitService := service.NewVisitService(repository)
	articleController := controllers.NewArticleController(articleService, visitService)

	eArticles := e.Group("/articles")
	eArticles.GET("", articleController.GetArticlesController)
	eArticles.GET("/:articleId", articleController.GetArticleController)

	eAuth := eArticles.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))
	eAuth.POST("", articleController.AddArticleController)

	eAuthArticle := eAuth.Group("/:articleId")
	eAuthArticle.PUT("", articleController.UpdateArticleController)
	eAuthArticle.DELETE("", articleController.DeleteArticleController)

	LikeRoute(eAuthArticle)
	CommentRoute(eAuthArticle)
}

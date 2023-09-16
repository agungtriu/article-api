package routes

import (
	"article-api/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ArticleRoute(e *echo.Echo) {
	eArticles := e.Group("/articles")
	eArticles.GET("", controllers.GetArticlesController)
	eArticles.GET("/:articleId", controllers.GetArticleController)

	eAuth := eArticles.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.POST("", controllers.AddArticleController)

	eAuthArticle := eAuth.Group("/:articleId")
	eAuthArticle.PUT("", controllers.UpdateArticleController)
	eAuthArticle.DELETE("", controllers.DeleteArticleController)

	LikeRoute(eAuthArticle)
	CommentRoute(eAuthArticle)
}

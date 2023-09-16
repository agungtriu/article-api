package main

import (
	"article-api/configs"
	"article-api/helper"
	"article-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	helper.LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	e.Use(middleware.Logger())
	routes.UserRoute(e)
	routes.ArticleRoute(e)
	e.Start(helper.GetPort())
}

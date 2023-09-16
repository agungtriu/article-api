package routes

import (
	"article-api/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	eUsers := e.Group("/users")

	eUsers.POST("/register", controllers.RegisterController)
	eUsers.POST("/login", controllers.LoginController)
	eUsers.GET("/:userId", controllers.GetUserController)

	eAuth := eUsers.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.PUT("/username", controllers.ChangeUsernameController)
	eAuth.PUT("/password", controllers.ChangePasswordController)
	eAuth.PUT("/email", controllers.ChangeEmailController)
	eAuth.PUT("/profile", controllers.ChangeProfileController)
}

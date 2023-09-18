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

func UserRoute(e *echo.Echo) {
	repository := repository.NewRepository(configs.DB)
	userService := service.NewUserService(repository)
	profileService := service.NewProfileService(repository)
	userController := controllers.NewUserController(userService, profileService)
	profileController := controllers.NewProfileController(profileService)

	eUsers := e.Group("/users")

	eUsers.POST("/register", userController.RegisterController)
	eUsers.POST("/login", userController.LoginController)
	eUsers.GET("/:userId", userController.GetUserController)

	eAuth := eUsers.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))
	eAuth.PUT("/username", userController.ChangeUsernameController)
	eAuth.PUT("/password", userController.ChangePasswordController)
	eAuth.PUT("/email", userController.ChangeEmailController)
	eAuth.PUT("/profile", profileController.ChangeProfileController)
}

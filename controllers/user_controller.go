package controllers

import (
	"article-api/helper"
	"article-api/middlewares"
	"article-api/models/base"
	profileResponse "article-api/models/profile/response"
	"article-api/models/user/request"
	userResponse "article-api/models/user/response"
	"article-api/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService    service.UserService
	profileService service.ProfileService
}

func NewUserController(userService service.UserService, profileService service.ProfileService) *userController {
	return &userController{userService, profileService}
}

func (controller *userController) RegisterController(c echo.Context) error {
	var requestRegister request.Register

	c.Bind(&requestRegister)

	if requestRegister.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if strings.Contains(requestRegister.Username, " ") {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid username format",
		})
	}
	if requestRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email cannot be empty",
		})
	}
	if !helper.IsEmailValid(requestRegister.Email) {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid email format",
		})
	}

	if requestRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	if requestRegister.ConfirmPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Confirm Password cannot be empty",
		})
	}

	if requestRegister.Password != requestRegister.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password and Confirm Password not match",
		})
	}

	channelUser := make(chan userResponse.Result)
	go controller.userService.VerifyUsername(requestRegister.Username, channelUser)
	user := <-channelUser

	if user.Err == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username already registered",
		})
	}

	go controller.userService.VerifyEmail(requestRegister.Email, channelUser)
	user = <-channelUser

	if user.Err == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email already registered",
		})
	}

	go controller.userService.RegisterUser(requestRegister, channelUser)
	user = <-channelUser
	if user.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  user.Err.Error(),
		})

	}
	channelProfile := make(chan profileResponse.Result)
	go controller.profileService.RegisterProfile(int(user.User.ID), channelProfile)
	profile := <-channelProfile
	if profile.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  profile.Err.Error(),
		})

	}

	var responseRegister userResponse.Register
	responseRegister.MapRegisterFromDatabase(user.User)

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success register user",
		Data:    responseRegister,
	})

}

func (controller *userController) LoginController(c echo.Context) error {
	var requestLogin request.Login

	c.Bind(&requestLogin)

	if requestLogin.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if requestLogin.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	channelUser := make(chan userResponse.Result)
	go controller.userService.VerifyUsername(requestLogin.Username, channelUser)
	user := <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  requestLogin.Username + " was not registered",
		})
	}

	isErr := helper.CheckPasswordHash(requestLogin.Password, user.User.Password)

	if isErr {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	var responseLogin userResponse.Login
	responseLogin.MapLoginFromDatabase(user.User)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success login user",
		Data:    responseLogin,
	})

}

func (controller *userController) GetUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))

	channelUser := make(chan userResponse.Result)
	go controller.userService.GetUser(userId, channelUser)
	user := <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "User not found",
		})
	}

	var responseUser userResponse.User
	responseUser.MapUserFromDatabase(user.User)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data user",
		Data:    responseUser,
	})
}

func (controller *userController) ChangeUsernameController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestChangeUsername request.ChangeUsername
	c.Bind(&requestChangeUsername)

	if requestChangeUsername.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if strings.Contains(requestChangeUsername.Username, " ") {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot contain space",
		})
	}

	if requestChangeUsername.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	channelUser := make(chan userResponse.Result)
	go controller.userService.GetUser(userId, channelUser)
	user := <-channelUser

	isErr := helper.CheckPasswordHash(requestChangeUsername.Password, user.User.Password)

	if isErr {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	go controller.userService.VerifyUsername(requestChangeUsername.Username, channelUser)
	user = <-channelUser

	if user.Err == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username already registered",
		})
	}

	go controller.userService.ChangeUsername(userId, requestChangeUsername.Username, channelUser)
	user = <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: true,
			Error:  "Failed to change username",
		})
	}
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success change username",
	})

}

func (controller *userController) ChangePasswordController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestChangePassword request.ChangePassword
	c.Bind(&requestChangePassword)

	if requestChangePassword.OldPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Old Password cannot be empty",
		})
	}

	if requestChangePassword.NewPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "New Password cannot be empty",
		})
	}

	if requestChangePassword.ConfirmPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Confirm Password cannot be empty",
		})
	}

	if requestChangePassword.NewPassword != requestChangePassword.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password and Confirm Password not match",
		})
	}

	channelUser := make(chan userResponse.Result)
	go controller.userService.GetUser(userId, channelUser)
	user := <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "User not found",
		})
	}

	isErr := helper.CheckPasswordHash(requestChangePassword.OldPassword, user.User.Password)

	if isErr {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	go controller.userService.ChangePassword(userId, requestChangePassword.NewPassword, channelUser)
	user = <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: true,
			Error:  "Failed to change password",
		})
	}

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success change password",
	})

}

func (controller *userController) ChangeEmailController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestChangeEmail request.ChangeEmail
	c.Bind(&requestChangeEmail)

	if requestChangeEmail.Email == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email cannot be empty",
		})
	}

	if !helper.IsEmailValid(requestChangeEmail.Email) {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid email format",
		})
	}

	if requestChangeEmail.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	channelUser := make(chan userResponse.Result)
	go controller.userService.GetUser(userId, channelUser)
	user := <-channelUser

	isErr := helper.CheckPasswordHash(requestChangeEmail.Password, user.User.Password)

	if isErr {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	go controller.userService.VerifyEmail(requestChangeEmail.Email, channelUser)
	user = <-channelUser

	if user.Err == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email already registered",
		})
	}

	go controller.userService.ChangeEmail(userId, requestChangeEmail.Email, channelUser)
	user = <-channelUser

	if user.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: true,
			Error:  "Failed to change email",
		})
	}
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success change email",
	})

}

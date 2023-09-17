package controllers

import (
	"article-api/configs"
	"article-api/helper"
	"article-api/middlewares"
	"article-api/models/base"
	profiledatabase "article-api/models/profile/database"
	userdatabase "article-api/models/user/database"
	userRequest "article-api/models/user/request"
	"article-api/models/user/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterController(c echo.Context) error {
	var userRegister userRequest.UserRegister

	c.Bind(&userRegister)

	if userRegister.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if strings.Contains(userRegister.Username, " ") {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid username format",
		})
	}
	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email cannot be empty",
		})
	}
	if !helper.IsEmailValid(userRegister.Email) {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid email format",
		})
	}

	if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	if userRegister.ConfirmPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Confirm Password cannot be empty",
		})
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password and Confirm Password not match",
		})
	}

	var userDatabase userdatabase.User

	varifyUserEmail := configs.DB.First(&userDatabase, "username = ? OR email = ?", strings.ToLower(userRegister.Username), strings.ToLower(userRegister.Email))

	if varifyUserEmail.Error == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username or Email already registered",
		})
	}

	hashedPassword, _ := helper.HashPassword(userRegister.Password)

	user := userdatabase.User{Username: strings.ToLower(userRegister.Username), Email: strings.ToLower(userRegister.Email), Password: hashedPassword}
	result := configs.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})

	}

	profile := profiledatabase.Profile{Name: "", Bio: "", UserId: int(user.ID)}
	configs.DB.Create(&profile)

	var registerResponse response.RegisterResponse
	registerResponse.MapRegisterFromDatabase(user)

	return c.JSON(http.StatusCreated, base.DataResponse{
		Status:  true,
		Message: "Success register user",
		Data:    registerResponse,
	})

}

func LoginController(c echo.Context) error {
	var userLogin userRequest.UserLogin

	c.Bind(&userLogin)

	if userLogin.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if userLogin.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	var userDatabase userdatabase.User
	key := strings.ToLower(userLogin.Username)
	verifyUserEmail := configs.DB.First(&userDatabase, "username = ? OR email = ?", key, key)

	if errors.Is(verifyUserEmail.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  userLogin.Username + " was not registered",
		})
	}

	err := helper.CheckPasswordHash(userLogin.Password, userDatabase.Password)

	if err {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}
	var loginResponse response.LoginResponse
	loginResponse.MapLoginFromDatabase(userDatabase)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success login user",
		Data:    loginResponse,
	})

}

func GetUserController(c echo.Context) error {
	userId := c.Param("userId")

	var user userdatabase.User
	err := configs.DB.Model(&userdatabase.User{}).Preload("Profile").First(&user, "id = ?", userId).Error

	if err != nil {
		return c.JSON(http.StatusNotFound, base.ErrorResponse{
			Status: false,
			Error:  "user was not found",
		})
	}

	var userResponse response.UserResponse
	userResponse.MapUserFromDatabase(user)

	return c.JSON(http.StatusOK, base.DataResponse{
		Status:  true,
		Message: "Success get data user",
		Data:    userResponse,
	})
}

func ChangeUsernameController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId := claims["userId"]

	var userChangeUsername userRequest.UserChangeUsername
	c.Bind(&userChangeUsername)

	if userChangeUsername.Username == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot be empty",
		})
	}

	if strings.Contains(userChangeUsername.Username, " ") {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username cannot contain space",
		})
	}

	if userChangeUsername.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	var userDatabase userdatabase.User
	configs.DB.First(&userDatabase, "id = ?", userId)

	err := helper.CheckPasswordHash(userChangeUsername.Password, userDatabase.Password)

	if err {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	varifyUsername := configs.DB.First(&userDatabase, "username = ?", strings.ToLower(userChangeUsername.Username))

	if varifyUsername.Error == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Username already registered",
		})
	}

	result := configs.DB.Model(&userDatabase).Where("id = ?", userId).Update("username", strings.ToLower(userChangeUsername.Username))

	if result.Error != nil {
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

func ChangePasswordController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId := claims["userId"]

	var userChangePassword userRequest.UserChangePassword
	c.Bind(&userChangePassword)

	if userChangePassword.OldPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Old Password cannot be empty",
		})
	}

	if userChangePassword.NewPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "New Password cannot be empty",
		})
	}

	if userChangePassword.ConfirmPassword == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Confirm Password cannot be empty",
		})
	}

	if userChangePassword.NewPassword != userChangePassword.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password and Confirm Password not match",
		})
	}
	var userDatabase userdatabase.User
	configs.DB.First(&userDatabase, "id = ?", userId)

	err := helper.CheckPasswordHash(userChangePassword.OldPassword, userDatabase.Password)

	if err {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}
	hashedPassword, _ := helper.HashPassword(userChangePassword.NewPassword)
	result := configs.DB.Model(&userDatabase).Where("id = ?", userId).Update("password", hashedPassword)

	if result.Error != nil {
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

func ChangeEmailController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId := claims["userId"]

	var userChangeEmail userRequest.UserChangeEmail
	c.Bind(&userChangeEmail)

	if userChangeEmail.Email == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email cannot be empty",
		})
	}

	if !helper.IsEmailValid(userChangeEmail.Email) {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Invalid email format",
		})
	}

	if userChangeEmail.Password == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Password cannot be empty",
		})
	}

	var userDatabase userdatabase.User
	configs.DB.First(&userDatabase, "id = ?", userId)

	err := helper.CheckPasswordHash(userChangeEmail.Password, userDatabase.Password)

	if err {
		return c.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Status: false,
			Error:  "Invalid password",
		})
	}

	varifyEmail := configs.DB.First(&userDatabase, "email = ?", strings.ToLower(userChangeEmail.Email))

	if varifyEmail.Error == nil {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Email already registered",
		})
	}

	result := configs.DB.Model(&userDatabase).Where("id = ?", userId).Update("email", strings.ToLower(userChangeEmail.Email))

	if result.Error != nil {
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

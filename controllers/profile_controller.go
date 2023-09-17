package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/base"
	"article-api/models/profile/request"
	"article-api/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ChangeProfileController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))

	var requestChangeProfile request.ChangeProfile
	c.Bind(&requestChangeProfile)

	if requestChangeProfile.Name == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Name cannot be empty",
		})
	}
	repository := repository.NewRepository(configs.DB)
	_, err := repository.ChangeProfile(userId, requestChangeProfile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success update profile",
	})

}

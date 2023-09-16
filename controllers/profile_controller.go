package controllers

import (
	"article-api/configs"
	"article-api/middlewares"
	"article-api/models/base"
	profiledatabase "article-api/models/profile/database"
	profileRequest "article-api/models/profile/request"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ChangeProfileController(c echo.Context) error {
	fullToken := c.Request().Header.Get("Authorization")
	token := strings.Split(fullToken, " ")
	claims, _ := middlewares.ExtractClaims(token[1])
	userId := claims["userId"]

	var changeProfile profileRequest.UserChangeProfile
	c.Bind(&changeProfile)

	if changeProfile.Name == "" {
		return c.JSON(http.StatusBadRequest, base.ErrorResponse{
			Status: false,
			Error:  "Name cannot be empty",
		})
	}
	var profileDatabase profiledatabase.Profile
	result := configs.DB.Model(&profileDatabase).Where("user_id = ?", userId).Updates(profiledatabase.Profile{Name: changeProfile.Name, Bio: changeProfile.Bio})

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  result.Error,
		})
	}
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success update profile",
	})

}

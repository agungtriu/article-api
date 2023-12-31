package controllers

import (
	"article-api/middlewares"
	"article-api/models/base"
	"article-api/models/profile/request"
	"article-api/models/profile/response"
	"article-api/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type profileController struct {
	profileService service.ProfileService
}

func NewProfileController(profileService service.ProfileService) *profileController {
	return &profileController{profileService}
}

func (controller *profileController) ChangeProfileController(c echo.Context) error {
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
	channelProfile := make(chan response.Result)
	go controller.profileService.ChangeProfile(userId, requestChangeProfile, channelProfile)
	profile := <-channelProfile

	if profile.Err != nil {
		return c.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Status: false,
			Error:  profile.Err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success update profile",
	})

}

package controllers

import (
	"miniproject/configs"
	"miniproject/models"
	"miniproject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []models.User
	if err := configs.DB.Preload("Membership").Find(&users).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}

	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! Get users",
		Data:    userResponses,
	}
	return c.JSON(http.StatusOK, response)

}

func GetUserController(c echo.Context) error {
	var user models.User
	if err := configs.DB.Preload("Membership").First(&user, c.Param("id")).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	userResponse := user.ToUserResponse()
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "success get user",
		Data:    userResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateUserController(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Create(user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := configs.DB.Preload("Membership").First(user, user.ID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "User not found")
	}
	userResponse := user.ToUserResponse()

	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! User created",
		Data:    userResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if err := configs.DB.First(&user, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "user not found")

	}

	if err := configs.DB.Delete(&user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "success Delete user",
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	if err := configs.DB.Preload("Membership").First(&user, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "user not found")

	}
	if err := c.Bind(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := configs.DB.Save(&user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	userResponse := user.ToUserResponse()
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "success Create user",
		Data:    userResponse,
	}
	return c.JSON(http.StatusOK, response)
}

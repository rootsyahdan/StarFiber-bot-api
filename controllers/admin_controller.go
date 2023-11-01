package controllers

import (
	"miniproject/configs"
	"miniproject/constants"
	"miniproject/middleware"
	"miniproject/models"
	"miniproject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateAdminController(c echo.Context) error {
	admin := new(models.Admin)
	if err := c.Bind(admin); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Encrypt password before saving
	err := middleware.EncryptAdminPassword(admin)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Password encryption failed")
	}

	if err := configs.DB.Create(admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Response for successful admin creation
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! created Admin",
		},
		Results: admin,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAdminsController(c echo.Context) error {
	var admin []models.Admin
	if err := configs.DB.Find(&admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Retrieved Admin",
		},
		Results: admin,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAdminController(c echo.Context) error {
	var admin models.Admin
	if err := configs.DB.First(&admin, c.Param("id")).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Get Admin",
		},
		Results: admin,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteAdminController(c echo.Context) error {
	id := c.Param("id")
	var admin models.Admin
	if err := configs.DB.First(&admin, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Admin not found")
	}
	if err := configs.DB.Delete(&admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Deleted Admin",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateAdminController(c echo.Context) error {
	id := c.Param("id")
	admin := models.Admin{}
	if err := configs.DB.First(&admin, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Admin not found")
	}
	if err := c.Bind(&admin); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := configs.DB.Save(&admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Updated Admin",
		},
		Results: admin,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginAdminController(c echo.Context) error {
	inputAdmin := new(models.Admin)
	if err := c.Bind(inputAdmin); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid data")
	}

	admin := new(models.Admin)

	if err := configs.DB.Where("username = ?", inputAdmin.Username).First(admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
	}

	if !middleware.ComparePasswords(admin.Password, inputAdmin.Password) {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := middleware.CreateToken(int(admin.ID), admin.Username, constants.SECRET_JWT)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Token creation failed")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Success": true,
		"message": "Success! Login successful",
		"Result":  admin,
		"token":   token,
	})
}

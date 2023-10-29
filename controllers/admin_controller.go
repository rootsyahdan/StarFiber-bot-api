package controllers

import (
	"miniproject/configs"
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

	if err := configs.DB.Create(admin).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := configs.DB.First(admin, admin.ID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Admin not found")
	}
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! User created",
		Data:    admin,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)
	err := configs.DB.Where("username = ? AND password = ? ", admin.Username, admin.Password).First(&admin).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
	}

	secretKey := "your-secret-key"
	token, err := middleware.CreateToken(int(admin.ID), admin.Username, secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Token creation failed",
			"error":   err.Error(),
		})
	}

	AdminResponse := models.AdminResponse{
		ID:       admin.ID,
		Username: admin.Username,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
		"data":    AdminResponse,
	})
}

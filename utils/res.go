package utils

import "github.com/labstack/echo/v4"

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	response := map[string]interface{}{
		"success": false,
		"message": message,
	}
	return c.JSON(statusCode, response)
}

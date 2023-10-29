package utils

import "github.com/labstack/echo/v4"

type JSONResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
		"data":    nil,
	}
	return c.JSON(statusCode, response)
}

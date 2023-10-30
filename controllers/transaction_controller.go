package controllers

import (
	"miniproject/configs"
	"miniproject/models"
	"miniproject/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTransactionsController(c echo.Context) error {
	requestedMonth := c.QueryParam("month")
	requestedYear := c.QueryParam("year")
	requestedStatus := c.QueryParam("status")

	var transactions []models.Transaction
	db := configs.DB.Model(&models.Transaction{}).Preload("User").Preload("User.Membership")

	if requestedMonth != "" && requestedYear != "" {
		month, err := strconv.Atoi(requestedMonth)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid 'month' query parameter.")
		}

		year, err := strconv.Atoi(requestedYear)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid 'year' query parameter.")
		}

		db = db.Where("MONTH(transaction_date) = ? AND YEAR(transaction_date) = ?", month, year)
	}

	if requestedStatus != "" {
		status, err := strconv.ParseBool(requestedStatus)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid 'status' query parameter. Use 'true' or 'false'.")
		}
		db = db.Where("status = ?", status)
	}

	if err := db.Find(&transactions).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var transactionResponses []models.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToTransactionResponse())
	}

	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! Retrieved transactions",
		Data:    transactionResponses,
	}

	return c.JSON(http.StatusOK, response)
}

func GetTransactionController(c echo.Context) error {
	var transaction models.Transaction
	if err := configs.DB.Preload("User.Membership").First(&transaction, c.Param("id")).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	transactionResponse := transaction.ToTransactionResponse()
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "success get transaction",
		Data:    transactionResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateTransactionController(c echo.Context) error {
	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Create(transaction).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := configs.DB.Preload("User.Membership").First(transaction, transaction.ID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Transaction not found")
	}

	transactionResponse := transaction.ToTransactionResponse()

	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! Create transaction",
		Data:    transactionResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteTransactionController(c echo.Context) error {
	id := c.Param("id")
	var transaction models.Transaction
	if err := configs.DB.First(&transaction, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Transaction not Found")
	}

	if err := configs.DB.Delete(&transaction).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "success Delete transaction",
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateTransactionController(c echo.Context) error {
	id := c.Param("id")
	transaction := models.Transaction{}

	if err := configs.DB.First(&transaction, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Transaction not Found")
	}

	if err := c.Bind(&transaction); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Save(&transaction).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if err := configs.DB.Preload("User.Membership").First(&transaction, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Transaction not Found")
	}

	transactionResponse := transaction.ToTransactionResponse()
	response := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Success! Update transaction",
		Data:    transactionResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateTransactionAutomaticallyController(c echo.Context) error {
	lastExecution := models.LastExecution{}
	if err := configs.DB.Last(&lastExecution).Error; err != nil && err != gorm.ErrRecordNotFound {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	now := time.Now()
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	if lastExecution.LastExec.Year() != currentMonth.Year() || lastExecution.LastExec.Month() != currentMonth.Month() {
		var users []models.User
		if err := configs.DB.Preload("Membership").Find(&users).Error; err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		for _, user := range users {
			newTransaction := models.Transaction{
				Status:          false,
				TransactionDate: time.Now(),
				UserID:          user.ID,
			}

			if err := configs.DB.Create(&newTransaction).Error; err != nil {
				return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			}

			requestNoInvoice := "INV" + strconv.FormatUint(uint64(newTransaction.ID), 10)
			newTransaction.NoInvoice = requestNoInvoice

			if err := configs.DB.Save(&newTransaction).Error; err != nil {
				return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			}
		}

		lastExecution.LastExec = currentMonth
		if err := configs.DB.Save(&lastExecution).Error; err != nil {
			return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		response := utils.JSONResponse{
			Status:  http.StatusCreated,
			Message: "success create transactions automatically",
			Data:    users,
		}
		return c.JSON(http.StatusCreated, response)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Transaction creation already executed for this month",
	})
}

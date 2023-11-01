package controllers

import (
	"context"
	"fmt"
	"miniproject/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type personRequest struct {
	Person int `json:"person"`
}

var client *openai.Client

func CreatePersonRequest(c echo.Context) error {
	var req personRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	message := fmt.Sprintf("Saran speed internet dalam mbps untuk %d orang ", req.Person)
	apiKey := os.Getenv("KEY")
	client = openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "Kamu Adalah Customer Services",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	data := resp.Choices[0].Message.Content

	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Get Recommendation",
		},
		Results: data,
	}

	return c.JSON(http.StatusOK, response)

}

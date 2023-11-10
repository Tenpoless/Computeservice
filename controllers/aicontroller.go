package controllers

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

func RecommendationAI(c echo.Context) error {
	userInput := c.FormValue("Gol_Darah")

	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	client := openai.NewClient(os.Getnev("AI_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Give a laptop recommendation" + userInput,
				},
			},
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menguhubngi AI"))
	}

	recommendationLaptop := resp.Choices[0].Message.Content

	return c.JSON(http.StatusOK, utils.SuccessResponse("Recommendation: ", recommendationLaptop))
}
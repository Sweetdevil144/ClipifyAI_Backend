package user

import (
	"ClipifyAI/config"
	"ClipifyAI/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type KeyInput struct {
	Username  string `json:"username"`
	OpenAIKey string `json:"openai_key,omitempty"`
	ClaudeKey string `json:"claude_key,omitempty"`
	YTKey     string `json:"yt_key,omitempty"`
	GeminiKey string `json:"gemeni_key,omitempty"`
}

func UpdateKeys(c *fiber.Ctx) error {
	var input KeyInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if input.OpenAIKey == "" && input.ClaudeKey == "" && input.YTKey == "" && input.GeminiKey == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "At least one key is required",
		})
	}

	user := models.User{}
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if input.OpenAIKey != "" {
		user.OpenAIApiKey = &input.OpenAIKey
	}
	if input.ClaudeKey != "" {
		user.ClaudeApiKey = &input.ClaudeKey
	}
	if input.YTKey != "" {
		user.YTApiKey = &input.YTKey
	}
	if input.GeminiKey != "" {
		user.GeminiApiKey = &input.GeminiKey
	}

	// Save updated user
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update keys",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Keys updated successfully",
	})
}

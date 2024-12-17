package user

import (
	"ClipifyAI/config"
	"ClipifyAI/models"
	"ClipifyAI/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	// Parse Input
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Input Validation
	if input.Username == "" || input.Name == "" || input.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "All fields are required",
		})
	}
	if !utils.IsSafe(input.Password) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Password is not strong enough",
		})
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create User Instance
	user := models.User{
		Username: input.Username,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	// Save User in DB
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate JWT Token
	token, err := utils.SerialiseUser(input.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"token":   token,
	})
}

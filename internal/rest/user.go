package rest

import (
	"goserver/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserRequest struct for capturing user creation data from request body
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"email"`
}

// CreateUserResponse struct for responding with user creation status and data
type CreateUserResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"` // include data only on success
}

func createUser(ctx *fiber.Ctx) error {

	var input CreateUserRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(CreateUserResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
		})
	}

	// 2. Validate user data
	if len(input.Username) < 3 {
		return ctx.Status(fiber.StatusBadRequest).JSON(CreateUserResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Username must be at least 3 characters long",
		})
	}

	// 3. Hash password securely using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(CreateUserResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error hashing password",
		})
	}

	// 4. Create a new User struct
	user := models.User{
		Username:  input.Username,
		Password:  string(hashedPassword),
		Email:     input.Email,
		CreatedAt: time.Now(),
	}

	err = database.db.createUser(&user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(CreateUserResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error creating user",
		})
	}

	// 6. Respond with success message (excluding password from response)
	return ctx.Status(fiber.StatusBadRequest).JSON(CreateUserResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "User created successfully",
		Data:       map[string]interface{}{"id": user.ID, "username": user.Username, "email": user.Email},
	})
}

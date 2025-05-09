package handler

import (
	"SiamOutlet/internal/auth/domain"
	"SiamOutlet/internal/auth/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) AuthRoutes(app *fiber.App) {
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := h.uc.Register(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	token, err := h.uc.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	// เก็บ token ใน cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Template_JWT",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(fiber.Map{
		"message": "Login success",
	})
}

// Handler สำหรับ admin dashboard
func (h *AuthHandler) AdminDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to admin dashboard",
	})
}

// Handler สำหรับ user profile
func (h *AuthHandler) UserProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to user profile",
	})
}

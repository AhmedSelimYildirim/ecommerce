package http

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userUC *usecase.UserUsecase
}

func NewAuthHandler(userUC *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{userUC: userUC}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	// TODO: implement user registration
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	// TODO: implement login
	return c.SendStatus(fiber.StatusNotImplemented)
}

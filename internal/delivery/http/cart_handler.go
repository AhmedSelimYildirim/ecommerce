package http

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	usecase *usecase.CartUsecase
}

func NewCartHandler(uc *usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase: uc}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *CartHandler) UpdateCart(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

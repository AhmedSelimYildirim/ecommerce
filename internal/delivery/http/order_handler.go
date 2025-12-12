package http

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(uc *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	// TODO: implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

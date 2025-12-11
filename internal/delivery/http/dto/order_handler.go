package http

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(uc *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

// CreateOrder endpoint
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	type Request struct {
		ProductIDs []int `json:"product_ids"`
	}

	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// userID'yi JWT veya c.Locals() üzerinden al
	userIDStr := c.Locals("user_id").(string)
	userID, _ := strconv.Atoi(userIDStr)

	if err := h.usecase.CreateOrder(userID, req.ProductIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "order created"})
}

// GetUserOrders endpoint
func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, _ := strconv.Atoi(userIDStr)

	orders, err := h.usecase.GetUserOrders(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orders)
}

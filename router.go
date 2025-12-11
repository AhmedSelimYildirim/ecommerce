package main

import (
	"github.com/AhmedSelimYildirim/ecommerce/config"
	"github.com/AhmedSelimYildirim/ecommerce/internal/delivery/http"
	"github.com/AhmedSelimYildirim/ecommerce/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(
	app *fiber.App,
	cfg *config.Config,
	userHandler *http.UserHandler,
	productHandler *http.ProductHandler,
	cartHandler *http.CartHandler,
	orderHandler *http.OrderHandler,
) {
	api := app.Group("/api")

	// User routes
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Protected User routes
	userRoutes := api.Group("/users", middleware.JWTMiddleware(cfg.JWTSecret))
	userRoutes.Get("/:id", userHandler.GetUser)

	// Cart routes
	cartRoutes := api.Group("/cart", middleware.JWTMiddleware(cfg.JWTSecret))
	cartRoutes.Post("/", cartHandler.AddToCart)
	cartRoutes.Get("/", cartHandler.GetCart)
	cartRoutes.Put("/:id", cartHandler.UpdateCart)
	cartRoutes.Delete("/:id", cartHandler.RemoveFromCart)

	// Order routes
	orderRoutes := api.Group("/orders", middleware.JWTMiddleware(cfg.JWTSecret))
	orderRoutes.Get("/", orderHandler.GetUserOrders) // Düzeltilmiş fonksiyon ismi
	orderRoutes.Post("/", orderHandler.CreateOrder)

	// Product routes
	api.Get("/products", productHandler.GetAllProducts)
	api.Get("/products/:id", productHandler.GetProductByID)

	// Admin protected product routes
	adminRoutes := api.Group("/products", middleware.JWTMiddleware(cfg.JWTSecret))
	adminRoutes.Post("/", func(c *fiber.Ctx) error {
		if c.Locals("role") != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "admin only"})
		}
		return productHandler.CreateProduct(c)
	})
	adminRoutes.Put("/:id", func(c *fiber.Ctx) error {
		if c.Locals("role") != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "admin only"})
		}
		return productHandler.UpdateProduct(c)
	})
	adminRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		if c.Locals("role") != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "admin only"})
		}
		return productHandler.DeleteProduct(c)
	})
}

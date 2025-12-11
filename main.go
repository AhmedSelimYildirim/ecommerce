package main

import (
	"log"

	"github.com/AhmedSelimYildirim/ecommerce/config"
	"github.com/AhmedSelimYildirim/ecommerce/internal/delivery/http"
	"github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"
	"github.com/AhmedSelimYildirim/ecommerce/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Config yükle
	cfg := config.LoadConfig()
	log.Println("Config yüklendi:", cfg.AppPort)

	// Fiber app başlat
	app := fiber.New()

	// DB bağlantısı
	db := pg.NewPostgresDB(cfg)

	// Repository oluştur
	userRepo := pg.NewUserRepo(db)
	productRepo := pg.NewProductRepo(db)
	cartRepo := pg.NewCartRepo(db)

	// Usecase oluştur
	userUC := usecase.NewUserUsecase(userRepo, cfg.JWTSecret, cfg.JWTExpireHours)
	productUC := usecase.NewProductUsecase(productRepo)
	cartUC := usecase.NewCartUsecase(cartRepo)

	// Handler oluştur
	userHandler := http.NewUserHandler(userUC)
	productHandler := http.NewProductHandler(productUC)
	cartHandler := http.NewCartHandler(cartUC)

	// Router’ları kur
	SetupRouter(app, cfg, userHandler, productHandler, cartHandler)

	// Server başlat
	log.Fatal(app.Listen(":" + cfg.AppPort))
}

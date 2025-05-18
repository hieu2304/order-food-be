package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/hieu2304/order-food-be/config"
	_ "github.com/hieu2304/order-food-be/docs"
	order_repo "github.com/hieu2304/order-food-be/repos/order"
	product_repo "github.com/hieu2304/order-food-be/repos/product"
	"github.com/hieu2304/order-food-be/routes/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Auto migrate database tables
	productRepo := product_repo.NewRepository()
	if err := productRepo.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate product tables: %v", err)
	}

	orderRepo := order_repo.NewRepository()
	if err := orderRepo.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate order tables: %v", err)
	}

	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiber.Use(logger.New())

	// Initialize Swagger
	fiber.Get("/swagger/*", swagger.HandlerDefault)

	app.RegisterRoutes(fiber)

	port := config.GetPort()
	log.Printf("🚀 Server running on http://localhost:%s", port)
	log.Fatal(fiber.Listen(":" + port))
}

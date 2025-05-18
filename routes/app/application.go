package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieu2304/order-food-be/routes/order"
	product "github.com/hieu2304/order-food-be/routes/product"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	order.RegisterOrderRoutes(api.Group("/order"))
	product.RegisterProductRoutes(api.Group("/product"))
}

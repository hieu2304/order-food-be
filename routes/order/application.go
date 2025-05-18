package order

import (
	"github.com/gofiber/fiber/v2"
	order_model "github.com/hieu2304/order-food-be/models/order"
	order_service "github.com/hieu2304/order-food-be/services/order"
)

type Application struct {
	service *order_service.Service
}

func NewApplication() *Application {
	return &Application{
		service: order_service.NewService(),
	}
}

// @Summary Create a new order
// @Description Create a new order with the provided items
// @Tags orders
// @Accept json
// @Produce json
// @Param order body order_model.OrderRequest true "Order Request"
// @Success 200 {object} order_model.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /order [post]

func (a *Application) CreateOrder(c *fiber.Ctx) error {
	var req order_model.OrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	order, err := a.service.CreateOrder(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}
	return c.JSON(order)
}

// @Summary Get an order by ID
// @Description Get an order by its ID
// @Tags orders
// @Accept json
// @Produce json

func (a *Application) GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := a.service.GetOrder(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get order",
		})
	}
	return c.JSON(order)
}

func RegisterOrderRoutes(orderRouter fiber.Router) {
	orderRouter.Post("/", NewApplication().CreateOrder)
	orderRouter.Get("/:id", NewApplication().GetOrder)
}

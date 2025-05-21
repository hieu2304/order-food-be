package product

import (
	"github.com/gofiber/fiber/v2"
	product_model "github.com/hieu2304/order-food-be/models/product"
	product_service "github.com/hieu2304/order-food-be/services/product"
)

type Application struct {
	service *product_service.Service
}

func NewApplication() *Application {
	return &Application{
		service: product_service.NewService(),
	}
}

// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {array} product_model.Product
// @Router /product [get]
func (a *Application) GetAll(c *fiber.Ctx) error {
	pagination := &product_model.Pagination{
		Page:     c.QueryInt("page", 1),
		PageSize: c.QueryInt("page_size", 10),
	}
	products, err := a.service.GetAll(pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}
	return c.JSON(products)
}

// @Summary Get product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} product_model.Product
// @Failure 404 {object} map[string]string
// @Router /product/{id} [get]
func (a *Application) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := a.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch product",
		})
	}
	return c.JSON(product)
}

// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body product_model.Product true "Product"
// @Success 200 {object} product_model.Product
// @Failure 400 {object} map[string]string
// @Router /product [post]
func (a *Application) Create(c *fiber.Ctx) error {
	var product product_model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := a.service.Create(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

// @Summary Update a product
// @Description Update a product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body product_model.Product true "Product"
// @Success 200 {object} product_model.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product/{id} [put]
func (a *Application) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var product product_model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := a.service.Update(id, &product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}
	return c.JSON(product)
}

// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product/{id} [delete]
func (a *Application) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func RegisterProductRoutes(productRouter fiber.Router) {
	productRouter.Get("/", NewApplication().GetAll)
	productRouter.Get("/:id", NewApplication().GetByID)
	productRouter.Post("/", NewApplication().Create)
	productRouter.Put("/:id", NewApplication().Update)
	productRouter.Delete("/:id", NewApplication().Delete)
}

package order

import (
	"errors"
	"strconv"

	order_model "github.com/hieu2304/order-food-be/models/order"
	order_repo "github.com/hieu2304/order-food-be/repos/order"
	product_repo "github.com/hieu2304/order-food-be/repos/product"
)

type Service struct {
	repo        *order_repo.Repository
	productRepo *product_repo.Repository
}

func NewService() *Service {
	return &Service{
		repo:        order_repo.NewRepository(),
		productRepo: product_repo.NewRepository(),
	}
}

func (s *Service) CreateOrder(req *order_model.OrderRequest) (*order_model.Order, error) {
	if err := s.validateOrderRequest(req); err != nil {
		return nil, err
	}
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}

	orderItems := make([]order_model.OrderItem, len(req.Items))
	for i, item := range req.Items {
		productExists := false
		for _, p := range products {
			if p.ID == item.ProductID {
				productExists = true
				break
			}
		}
		if !productExists {
			return nil, errors.New("product not found: " + strconv.FormatUint(uint64(item.ProductID), 10))
		}

		orderItems[i] = order_model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	order := &order_model.Order{
		Items:    orderItems,
		Products: products,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, errors.New("failed to create order")
	}

	return order, nil
}

func (s *Service) GetOrder(id string) (*order_model.Order, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	order, err := s.repo.FindByID(uint(idUint))
	if err != nil {
		return nil, errors.New("failed to get order")
	}
	return order, nil
}
func (s *Service) validateOrderRequest(req *order_model.OrderRequest) error {
	if len(req.Items) == 0 {
		return errors.New("order must contain at least one item")
	}

	for _, item := range req.Items {
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
		if item.ProductID == 0 {
			return errors.New("product ID is required")
		}
	}

	return nil
}

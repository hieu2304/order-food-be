package product_service

import (
	"strconv"

	product_model "github.com/hieu2304/order-food-be/models/product"
	product_repo "github.com/hieu2304/order-food-be/repos/product"
)

type Service struct {
	repo *product_repo.Repository
}

func NewService() *Service {
	return &Service{
		repo: product_repo.NewRepository(),
	}
}

func (s *Service) GetAll(pagination *product_model.Pagination) ([]product_model.Product, error) {
	return s.repo.FindAll(pagination)
}

func (s *Service) GetByID(id string) (*product_model.Product, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(uint(idUint))
}

func (s *Service) Create(product *product_model.Product) error {
	return s.repo.Create(product)
}

func (s *Service) Update(id string, product *product_model.Product) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = s.repo.FindByID(uint(idUint))
	if err != nil {
		return err
	}
	return s.repo.Update(product)
}

func (s *Service) Delete(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return s.repo.Delete(uint(idUint))
}

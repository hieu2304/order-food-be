package product_repo

import (
	"github.com/hieu2304/order-food-be/config"
	product_model "github.com/hieu2304/order-food-be/models/product"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{
		db: config.DB,
	}
}

func (r *Repository) Create(product *product_model.Product) error {
	return r.db.Create(product).Error
}

func (r *Repository) FindAll() ([]product_model.Product, error) {
	var products []product_model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *Repository) FindByID(id uint) (*product_model.Product, error) {
	var product product_model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) Update(product *product_model.Product) error {
	return r.db.Save(product).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&product_model.Product{}, id).Error
}

func (r *Repository) AutoMigrate() error {
	return r.db.AutoMigrate(&product_model.Product{})
}

package order

import (
	"github.com/hieu2304/order-food-be/config"
	model_order "github.com/hieu2304/order-food-be/models/order"
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

func (r *Repository) Create(order *model_order.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Items").Create(order).Error; err != nil {
			return err
		}

		for i := range order.Items {
			order.Items[i].OrderID = order.ID
			if err := tx.Create(&order.Items[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *Repository) FindByID(id uint) (*model_order.Order, error) {
	var order model_order.Order
	if err := r.db.Preload("Items").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) AutoMigrate() error {
	if err := r.db.Migrator().DropTable(&model_order.OrderItem{}, &model_order.Order{}); err != nil {
		return err
	}

	return r.db.AutoMigrate(&model_order.Order{}, &model_order.OrderItem{})
}

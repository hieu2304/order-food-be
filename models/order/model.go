package model_order

import (
	product_model "github.com/hieu2304/order-food-be/models/product"
)

type Order struct {
	ID       uint                    `json:"-" gorm:"primaryKey;autoIncrement"`
	Items    []OrderItem             `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Products []product_model.Product `json:"products" gorm:"-"`
}

type OrderItem struct {
	ID        uint `json:"-" gorm:"primaryKey;autoIncrement"`
	OrderID   uint
	ProductID uint `json:"productId" gorm:"not null"`
	Quantity  int  `json:"quantity" gorm:"not null"`
}

type OrderRequest struct {
	Items []OrderItemRequest `json:"items" validate:"required,dive"`
}

type OrderItemRequest struct {
	ProductID uint `json:"productId" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

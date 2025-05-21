package product_service

import (
	"strconv"
	"testing"

	product_model "github.com/hieu2304/order-food-be/models/product"
	"github.com/hieu2304/order-food-be/utils"
	"github.com/stretchr/testify/assert"
)

func TestProductService_Create(t *testing.T) {
	utils.SetupTest(t)
	service := NewService()

	tests := []struct {
		name    string
		product *product_model.Product
		wantErr bool
	}{
		{
			name: "Product 1",
			product: &product_model.Product{
				Name:     "Pizza",
				Price:    10.99,
				Category: "Food",
			},
			wantErr: false,
		},
		{
			name: "invalid product - empty name",
			product: &product_model.Product{
				Name:     "",
				Price:    10.99,
				Category: "Food",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Create(tt.product)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.product.ID)
			}
		})
	}
}

func TestProductService_GetAll(t *testing.T) {
	service := NewService()
	products := []*product_model.Product{
		{Name: "Pizza", Price: 10.99, Category: "Food"},
		{Name: "Burger", Price: 5.99, Category: "Food"},
	}

	for _, p := range products {
		err := service.Create(p)
		assert.NoError(t, err)
	}

	pagination := &product_model.Pagination{
		Page:     1,
		PageSize: 10,
	}
	result, err := service.GetAll(pagination)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(result), len(products))
}

func TestProductService_GetByID(t *testing.T) {
	service := NewService()

	product := &product_model.Product{
		Name:     "Pizza",
		Price:    10.99,
		Category: "Food",
	}
	err := service.Create(product)
	assert.NoError(t, err)

	result, err := service.GetByID(strconv.FormatUint(uint64(product.ID), 10))
	assert.NoError(t, err)
	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Price, result.Price)
	assert.Equal(t, product.Category, result.Category)
}

package order

import (
	// Import path/filepath

	"testing"

	order_model "github.com/hieu2304/order-food-be/models/order"
	product_model "github.com/hieu2304/order-food-be/models/product"
	product_service "github.com/hieu2304/order-food-be/services/product"
	"github.com/hieu2304/order-food-be/utils"
	"github.com/stretchr/testify/assert"
)

func TestOrderService_CreateOrder(t *testing.T) {
	utils.SetupTest(t)
	service := NewService()

	tests := []struct {
		name    string
		req     *order_model.OrderRequest
		wantErr bool
	}{
		{
			name: "valid order",
			req: &order_model.OrderRequest{
				Items: []order_model.OrderItemRequest{
					{
						ProductID: 1,
						Quantity:  2,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid order - empty items",
			req: &order_model.OrderRequest{
				Items: []order_model.OrderItemRequest{},
			},
			wantErr: true,
		},
		{
			name: "invalid order - zero quantity",
			req: &order_model.OrderRequest{
				Items: []order_model.OrderItemRequest{
					{
						ProductID: 1,
						Quantity:  0,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productService := product_service.NewService()
			productService.Create(&product_model.Product{
				ID:       1,
				Name:     "Test Product",
				Price:    10.0,
				Category: "Test Category",
			})

			order, err := service.CreateOrder(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, order)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.NotZero(t, order.ID)
				assert.Equal(t, len(tt.req.Items), len(order.Items))
			}
		})
	}
}

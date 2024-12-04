package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/ecommerce/internal/domain"
)

func TestOrderRepository(t *testing.T) {
	repo := NewOrderRepository()

	t.Run("CreateOrder", func(t *testing.T) {
		order := &domain.Order{
			ID:     "order1",
			UserID: "user1",
			Cart: domain.Cart{
				UserID: "user1",
				Items: []domain.Item{
					{
						ProductID: "product1",
						Name:      "Test Product",
						Price:     10.0,
						Quantity:  2,
					},
				},
			},
			TotalAmount:   20.0,
			DiscountCode:  "SAVE10",
			DiscountTotal: 2.0,
			FinalAmount:   18.0,
		}

		err := repo.CreateOrder(order)
		assert.NoError(t, err)
	})

	t.Run("GetUserOrderCount", func(t *testing.T) {
		userID := "user2"
		
		// Create multiple orders for the same user
		order1 := &domain.Order{ID: "order2", UserID: userID}
		order2 := &domain.Order{ID: "order3", UserID: userID}

		err := repo.CreateOrder(order1)
		assert.NoError(t, err)

		err = repo.CreateOrder(order2)
		assert.NoError(t, err)

		count := repo.GetUserOrderCount(userID)
		assert.Equal(t, 2, count)
	})

	t.Run("ResetUserOrderCount", func(t *testing.T) {
		userID := "user3"
		
		// Create an order
		order := &domain.Order{ID: "order4", UserID: userID}
		err := repo.CreateOrder(order)
		assert.NoError(t, err)

		// Reset order count
		repo.ResetUserOrderCount(userID)

		count := repo.GetUserOrderCount(userID)
		assert.Equal(t, 0, count)
	})

	t.Run("GetOrderStats", func(t *testing.T) {
		// Clear previous orders
		repo = NewOrderRepository()

		// Create multiple orders with different characteristics
		order1 := &domain.Order{
			ID:            "order5",
			UserID:        "user4",
			TotalAmount:   100.0,
			DiscountCode:  "SAVE10",
			DiscountTotal: 10.0,
			FinalAmount:   90.0,
			Cart: domain.Cart{
				Items: []domain.Item{
					{ProductID: "product1", Quantity: 2},
					{ProductID: "product2", Quantity: 1},
				},
			},
		}

		order2 := &domain.Order{
			ID:            "order6",
			UserID:        "user5",
			TotalAmount:   200.0,
			DiscountCode:  "SAVE20",
			DiscountTotal: 40.0,
			FinalAmount:   160.0,
			Cart: domain.Cart{
				Items: []domain.Item{
					{ProductID: "product3", Quantity: 3},
				},
			},
		}

		err := repo.CreateOrder(order1)
		assert.NoError(t, err)

		err = repo.CreateOrder(order2)
		assert.NoError(t, err)

		stats := repo.GetOrderStats()
		assert.Equal(t, 3, stats["total_items_purchased"])
		assert.Equal(t, float64(250.0), stats["total_purchase_amount"])
		assert.Equal(t, float64(50.0), stats["total_discount_amount"])
		
		discountCodes := stats["discount_codes"].([]string)
		assert.Contains(t, discountCodes, "SAVE10")
		assert.Contains(t, discountCodes, "SAVE20")
	})
}
package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/ecommerce/internal/domain"
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/repository"
)

func setupAdminServiceTest() (*AdminService, *repository.OrderRepository, *repository.DiscountRepository) {
	orderRepo := repository.NewOrderRepository()
	discountRepo := repository.NewDiscountRepository()
	adminService := NewAdminService(orderRepo, discountRepo)
	return adminService, orderRepo, discountRepo
}

func TestAdminService_GenerateDiscountCode(t *testing.T) {
	t.Run("GenerateDiscountCode_FirstEligibleOrder", func(t *testing.T) {
		adminService, orderRepo, discountRepo := setupAdminServiceTest()
		userID := "user1"

		// Create multiple orders to reach the nth order threshold
		for i := 0; i < 5; i++ {
			order := &domain.Order{
				ID:     "order" + string(rune(i)),
				UserID: userID,
			}
			err := orderRepo.CreateOrder(order)
			assert.NoError(t, err)
		}

		// Generate discount code for 5th order
		code, err := adminService.GenerateDiscountCode(userID, 5, 10.0)
		assert.NoError(t, err)
		assert.NotEmpty(t, code)

		// Verify discount code was created
		discountCode, err := discountRepo.GetDiscountCode(code)
		assert.NoError(t, err)
		assert.NotNil(t, discountCode)
		assert.Equal(t, 10.0, discountCode.Percentage)
		assert.False(t, discountCode.IsUsed)
	})

	t.Run("GenerateDiscountCode_NotEligible", func(t *testing.T) {
		adminService, _, _ := setupAdminServiceTest()
		userID := "user2"

		// Try to generate discount code without meeting the order count
		code, err := adminService.GenerateDiscountCode(userID, 5, 10.0)
		assert.Error(t, err)
		assert.Empty(t, code)
		assert.Equal(t, errors.ErrDiscountCodeNotAvailable, err)
	})
}

func TestAdminService_GetAdminStats(t *testing.T) {
	t.Run("GetAdminStats_NoOrders", func(t *testing.T) {
		adminService, _, _ := setupAdminServiceTest()

		stats := adminService.GetAdminStats()
		assert.NotNil(t, stats)
		assert.Equal(t, 0, stats["total_items_purchased"])
		assert.Equal(t, float64(0), stats["total_purchase_amount"])
		assert.Equal(t, float64(0), stats["total_discount_amount"])
		assert.Len(t, stats["discount_codes"].([]string), 0)
	})

	t.Run("GetAdminStats_MultipleOrders", func(t *testing.T) {
		adminService, orderRepo, _ := setupAdminServiceTest()

		// Create multiple orders
		orders := []*domain.Order{
			{
				ID:            "order1",
				UserID:        "user1",
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
			},
			{
				ID:            "order2",
				UserID:        "user2",
				TotalAmount:   200.0,
				DiscountCode:  "SAVE20",
				DiscountTotal: 40.0,
				FinalAmount:   160.0,
				Cart: domain.Cart{
					Items: []domain.Item{
						{ProductID: "product3", Quantity: 3},
					},
				},
			},
		}

		for _, order := range orders {
			err := orderRepo.CreateOrder(order)
			assert.NoError(t, err)
		}

		stats := adminService.GetAdminStats()
		assert.NotNil(t, stats)
		assert.Equal(t, 3, stats["total_items_purchased"])
		assert.Equal(t, float64(250.0), stats["total_purchase_amount"])
		assert.Equal(t, float64(50.0), stats["total_discount_amount"])

		discountCodes := stats["discount_codes"].([]string)
		assert.Contains(t, discountCodes, "SAVE10")
		assert.Contains(t, discountCodes, "SAVE20")
	})
}

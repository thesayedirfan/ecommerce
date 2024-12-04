package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/ecommerce/internal/domain"
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/repository"
)

func setupCartServiceTest() (*CartService, *repository.CartRepository, *repository.OrderRepository, *repository.DiscountRepository) {
	cartRepo := repository.NewCartRepository()
	orderRepo := repository.NewOrderRepository()
	discountRepo := repository.NewDiscountRepository()
	cartService := NewCartService(orderRepo, cartRepo, discountRepo)
	return cartService, cartRepo, orderRepo, discountRepo
}

func TestCartService_AddToCart(t *testing.T) {
	cartService, _, _, _ := setupCartServiceTest()

	t.Run("AddToCart_Success", func(t *testing.T) {
		userID := "user1"
		err := cartService.AddToCart(userID, "product1", "Test Product", 10.0, 2)
		assert.NoError(t, err)
	})

	t.Run("AddToCart_MultipleItems", func(t *testing.T) {
		userID := "user2"
		err := cartService.AddToCart(userID, "product1", "Test Product 1", 10.0, 2)
		assert.NoError(t, err)

		err = cartService.AddToCart(userID, "product2", "Test Product 2", 15.0, 1)
		assert.NoError(t, err)
	})
}

func TestCartService_Checkout(t *testing.T) {
	t.Run("Checkout_Success_NoDiscount", func(t *testing.T) {
		cartService, cartRepo, _, _ := setupCartServiceTest()

		userID := "user1"
		// Add items to cart
		err := cartRepo.AddToCart(userID, domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     10.0,
			Quantity:  2,
		})
		assert.NoError(t, err)

		// Checkout
		order, err := cartService.Checkout(userID, "")
		assert.NoError(t, err)
		assert.NotNil(t, order)

		// Verify order details
		assert.Equal(t, userID, order.UserID)
		assert.Equal(t, float64(20.0), order.TotalAmount)
		assert.Equal(t, float64(20.0), order.FinalAmount)
		assert.Equal(t, "", order.DiscountCode)

		// Verify cart is cleared
		_, err = cartRepo.GetCart(userID)
		assert.Error(t, err)
	})

	t.Run("Checkout_Success_WithDiscount", func(t *testing.T) {
		cartService, cartRepo, _, discountRepo := setupCartServiceTest()

		userID := "user2"
		// Create discount code
		discountCode := "SAVE10"
		err := discountRepo.CreateDiscountCode(discountCode, 10.0)
		assert.NoError(t, err)

		// Add items to cart
		err = cartRepo.AddToCart(userID, domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     100.0,
			Quantity:  1,
		})
		assert.NoError(t, err)

		// Checkout with discount
		order, err := cartService.Checkout(userID, discountCode)
		assert.NoError(t, err)
		assert.NotNil(t, order)

		// Verify order details
		assert.Equal(t, userID, order.UserID)
		assert.Equal(t, float64(100.0), order.TotalAmount)
		assert.Equal(t, float64(10.0), order.DiscountTotal)
		assert.Equal(t, float64(90.0), order.FinalAmount)
		assert.Equal(t, discountCode, order.DiscountCode)

		// Verify discount code is marked as used
		usedDiscount, err := discountRepo.GetDiscountCode(discountCode)
		assert.NoError(t, err)
		assert.True(t, usedDiscount.IsUsed)
		assert.Equal(t, userID, usedDiscount.UsedByUserID)
	})

	t.Run("Checkout_CartNotFound", func(t *testing.T) {
		cartService, _, _, _ := setupCartServiceTest()

		userID := "nonexistent"
		order, err := cartService.Checkout(userID, "")
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Equal(t, errors.ErrCartNotFound, err)
	})

	t.Run("Checkout_InvalidDiscountCode", func(t *testing.T) {
		cartService, cartRepo, _, _ := setupCartServiceTest()

		userID := "user3"
		// Add items to cart
		err := cartRepo.AddToCart(userID, domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     100.0,
			Quantity:  1,
		})
		assert.NoError(t, err)

		// Checkout with invalid discount code
		order, err := cartService.Checkout(userID, "INVALID_CODE")
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Equal(t, errors.ErrInvalidDiscountCode, err)
	})
}

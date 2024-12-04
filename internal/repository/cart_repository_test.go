package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/ecommerce/internal/domain"
)

func TestCartRepository(t *testing.T) {
	repo := NewCartRepository()

	t.Run("AddToCart_NewUser", func(t *testing.T) {
		userID := "user1"
		item := domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     10.0,
			Quantity:  2,
		}

		err := repo.AddToCart(userID, item)
		assert.NoError(t, err)

		cart, err := repo.GetCart(userID)
		assert.NoError(t, err)
		assert.Equal(t, userID, cart.UserID)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, item, cart.Items[0])
	})

	t.Run("AddToCart_ExistingUser_NewProduct", func(t *testing.T) {
		userID := "user2"
		item1 := domain.Item{
			ProductID: "product1",
			Name:      "Test Product 1",
			Price:     10.0,
			Quantity:  2,
		}
		item2 := domain.Item{
			ProductID: "product2",
			Name:      "Test Product 2",
			Price:     15.0,
			Quantity:  1,
		}

		err := repo.AddToCart(userID, item1)
		assert.NoError(t, err)

		err = repo.AddToCart(userID, item2)
		assert.NoError(t, err)

		cart, err := repo.GetCart(userID)
		assert.NoError(t, err)
		assert.Len(t, cart.Items, 2)
	})

	t.Run("AddToCart_ExistingUser_ExistingProduct", func(t *testing.T) {
		userID := "user3"
		item1 := domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     10.0,
			Quantity:  2,
		}
		item2 := domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     10.0,
			Quantity:  3,
		}

		err := repo.AddToCart(userID, item1)
		assert.NoError(t, err)

		err = repo.AddToCart(userID, item2)
		assert.NoError(t, err)

		cart, err := repo.GetCart(userID)
		assert.NoError(t, err)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, 5, cart.Items[0].Quantity)
	})

	t.Run("GetCart_NonExistentUser", func(t *testing.T) {
		cart, err := repo.GetCart("nonexistent")
		assert.Error(t, err)
		assert.Nil(t, cart)
	})

	t.Run("ClearCart", func(t *testing.T) {
		userID := "user4"
		item := domain.Item{
			ProductID: "product1",
			Name:      "Test Product",
			Price:     10.0,
			Quantity:  2,
		}

		err := repo.AddToCart(userID, item)
		assert.NoError(t, err)

		err = repo.ClearCart(userID)
		assert.NoError(t, err)

		cart, err := repo.GetCart(userID)
		assert.Error(t, err)
		assert.Nil(t, cart)
	})
}
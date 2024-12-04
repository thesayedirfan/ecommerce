package repository

import (
	"sync"

	"github.com/thesayedirfan/ecommerce/internal/domain"
	"github.com/thesayedirfan/ecommerce/internal/errors"
)

type CartRepository struct {
	mu    sync.Mutex
	carts map[string]*domain.Cart
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		carts: make(map[string]*domain.Cart),
	}
}

func (r *CartRepository) AddToCart(userID string, item domain.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cart, exists := r.carts[userID]; exists {
		for i, existing := range cart.Items {
			if existing.ProductID == item.ProductID {
				cart.Items[i].Quantity += item.Quantity
				return nil
			}
		}
		cart.Items = append(cart.Items, item)
	} else {
		r.carts[userID] = &domain.Cart{
			UserID: userID,
			Items:  []domain.Item{item},
		}
	}

	return nil
}

func (r *CartRepository) GetCart(userID string) (*domain.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart, exists := r.carts[userID]
	if !exists {
		return nil, errors.ErrCartNotFound
	}
	return cart, nil
}

func (r *CartRepository) ClearCart(userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.carts, userID)
	return nil
}

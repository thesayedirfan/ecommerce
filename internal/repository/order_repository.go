package repository

import (
	"sync"
	"github.com/thesayedirfan/ecommerce/internal/domain"
)

type OrderRepository struct {
	mu             sync.Mutex
	orders         map[string]*domain.Order
	userOrderCount map[string]int
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders:         make(map[string]*domain.Order),
		userOrderCount: make(map[string]int),
	}
}

func (r *OrderRepository) CreateOrder(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.ID] = order
	r.userOrderCount[order.UserID]++
	return nil
}

func (r *OrderRepository) GetUserOrderCount(userID string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.userOrderCount[userID]
}

func (r *OrderRepository) ResetUserOrderCount(userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userOrderCount[userID] = 0
}

func (r *OrderRepository) GetOrderStats() map[string]interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	totalItems := 0
	totalPurchaseAmount := 0.0
	totalDiscountAmount := 0.0
	discountCode := make([]string,0)

	for _, order := range r.orders {
		totalItems += len(order.Cart.Items)
		totalPurchaseAmount += order.FinalAmount
		totalDiscountAmount +=  order.DiscountTotal
		if order.DiscountCode != "" {
			discountCode = append(discountCode, order.DiscountCode)
		}
	}

	return map[string]interface{}{
		"total_items_purchased": totalItems,
		"total_purchase_amount": totalPurchaseAmount,
		"total_discount_amount":totalDiscountAmount,
		"discount_codes":discountCode,
	}
}

package repository

import (
	"sync"

	"github.com/thesayedirfan/ecommerce/internal/domain"
	"github.com/thesayedirfan/ecommerce/internal/errors"
)

type DiscountRepository struct {
	mu             sync.Mutex
	discountCodes  map[string]*domain.DiscountCode
}

func NewDiscountRepository() *DiscountRepository {
	return &DiscountRepository{
		discountCodes:  make(map[string]*domain.DiscountCode),
	}
}

func (r *DiscountRepository) CreateDiscountCode(code string, percentage float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.discountCodes[code] = &domain.DiscountCode{
		Code:       code,
		Percentage: percentage,
		IsUsed:     false,
	}
	return nil
}

func (r *DiscountRepository) GetDiscountCode(code string) (*domain.DiscountCode, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	discountCode, exists := r.discountCodes[code]
	if !exists {
		return nil, nil
	}
	return discountCode, nil
}

func (r *DiscountRepository) UseDiscountCode(code, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if discountCode, exists := r.discountCodes[code]; exists {
		if discountCode.IsUsed {
			return errors.ErrDiscountCodeAlreadyUsed
		}
		
		discountCode.IsUsed = true
		discountCode.UsedByUserID = userID
	}
	return nil
}
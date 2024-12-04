package service

import (
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/repository"
	"github.com/thesayedirfan/ecommerce/internal/utils"
)

type AdminService struct {
	orderRepo    *repository.OrderRepository
	discountRepo *repository.DiscountRepository
}

func NewAdminService(
	orderRepo *repository.OrderRepository,
	discountRepo *repository.DiscountRepository,
) *AdminService {
	return &AdminService{
		orderRepo:    orderRepo,
		discountRepo: discountRepo,
	}
}

func (s *AdminService) GenerateDiscountCode(userID string,nthOrder int ,discountPercentage float64) (string, error) {
	orderCount := s.orderRepo.GetUserOrderCount(userID)
	if utils.CheckIfUserApplicableForDiscount(orderCount,nthOrder) {
		code := utils.GenerateUniqueDiscountCode()
		err := s.discountRepo.CreateDiscountCode(code, discountPercentage)
		if err != nil {
			return "", err
		}
		return code, nil
	}
	return "", errors.ErrDiscountCodeNotAvailable
}

func (s *AdminService) GetAdminStats() map[string]interface{} {
	orderStats := s.orderRepo.GetOrderStats()
	return orderStats
}

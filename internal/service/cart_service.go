package service

import (
	"fmt"

	"github.com/thesayedirfan/ecommerce/internal/domain"
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/repository"
	"github.com/thesayedirfan/ecommerce/internal/utils"
	"github.com/thesayedirfan/ecommerce/pkg/uuid"
)

type CartService struct {
	cartRepo     *repository.CartRepository
	orderRepo    *repository.OrderRepository
	discountRepo *repository.DiscountRepository
}

func NewCartService(
	orderRepo *repository.OrderRepository,
	cartRepo *repository.CartRepository,
	discountRepo *repository.DiscountRepository,
) *CartService {
	return &CartService{
		orderRepo:    orderRepo,
		cartRepo:     cartRepo,
		discountRepo: discountRepo,
	}
}

func (s *CartService) AddToCart(userID string, productID, name string, price float64, quantity int) error {
	item := domain.Item{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
	}
	return s.cartRepo.AddToCart(userID, item)
}

func (s *CartService) Checkout(userID string, discountCode string) (*domain.Order, error) {

	cart, err := s.cartRepo.GetCart(userID)
	if err != nil || cart == nil {
		return nil, errors.ErrCartNotFound
	}

	totalAmount := 0.0
	for _, item := range cart.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	fmt.Println(discountCode)

	finalAmount := totalAmount
	
	discountTotal := 0.0
	if discountCode != "" {
		discount, err := s.discountRepo.GetDiscountCode(discountCode)
		if err != nil || discount == nil || discount.IsUsed {
			return nil, errors.ErrInvalidDiscountCode
		}

		discountTotal = utils.GetDiscountedPrice(totalAmount,discount.Percentage)
		finalAmount -= discountTotal

		err = s.discountRepo.UseDiscountCode(discountCode, userID)
		if err != nil {
			return nil, err
		}
		
		s.orderRepo.ResetUserOrderCount(userID)
		
	}

	order := &domain.Order{
		ID:            uuid.LongUUID(),
		UserID:        userID,
		Cart:          *cart,
		TotalAmount:   totalAmount,
		DiscountCode:  discountCode,
		DiscountTotal: discountTotal,
		FinalAmount : finalAmount, 
	}

	err = s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	err = s.cartRepo.ClearCart(userID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

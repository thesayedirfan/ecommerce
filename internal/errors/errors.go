package errors

import "errors"

var (
	ErrCartNotFound =  errors.New("cart not found")
	ErrInvalidDiscountCode = errors.New("invalid discount code")
	ErrNotEligibleForDiscountCode = errors.New("not eligible for discount code yet")	
	ErrDiscountCodeAlreadyUsed = errors.New("discount code is already used")
	ErrDiscountCodeNotAvailable = errors.New("discount code not available")
	ErrUserIDCannotBeEmpty = errors.New("user id cannot be empty")
)

var (
	ErrItemProductIDCannotBeEmpty = errors.New("item product id cannot be empty")
	ErrItemNameCannotBeEmpty = errors.New("item name cannot be empty")
	ErrItemPriceCannotBeZero = errors.New("item price cannot be empty")
	ErrItemQuantityCannotBeZero = errors.New("item quantity cannot be empty")
)
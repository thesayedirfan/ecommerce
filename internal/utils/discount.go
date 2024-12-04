package utils

import (
	"fmt"

	"github.com/thesayedirfan/ecommerce/pkg/uuid"
)

func GenerateUniqueDiscountCode() string {
	return fmt.Sprintf("DISCOUNT-%s", uuid.ShortUUID())
}

func GetDiscountedPrice(totalamount float64 , percentage float64) (float64) {
	return totalamount * (percentage / 100)
}

func CheckIfUserApplicableForDiscount(orderCount int , nthOrder int) bool {
	return orderCount > 0 && (orderCount%nthOrder) == 0
}
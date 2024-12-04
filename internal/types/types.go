package types

type GenerateDiscountRequest struct {
	UserID   string `json:"user_id"`
}

type AddToCartRequest struct {
	UserID     string  `json:"user_id"`
	ProductID  string  `json:"product_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}


type CheckoutRequest struct {
	UserID       string `json:"user_id"`
	DiscountCode string `json:"discount_code,omitempty"`
}
package domain

type Order struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	Cart          Cart    `json:"cart"`
	TotalAmount   float64 `json:"total_amount"`
	DiscountCode  string  `json:"discount_code,omitempty"`
	DiscountTotal float64 `json:"discount_total,omitempty"`
	FinalAmount float64    `json:"final_amount"`
}

package domain

type DiscountCode struct {
	Code         string  `json:"code"`
	Percentage   float64 `json:"percentage"`
	UsedByUserID string  `json:"used_by_user_id,omitempty"`
	IsUsed       bool    `json:"is_used"`
}
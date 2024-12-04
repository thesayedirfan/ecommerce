package domain

type Cart struct {
	UserID string `json:"user_id"`
	Items  []Item `json:"items"`
}

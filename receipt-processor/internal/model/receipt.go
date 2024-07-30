package model

type Receipt struct {
	ID           string  `json:"id,omitempty"`
	Retailer     string  `json:"retailer" binding:"required"`
	PurchaseDate string  `json:"purchaseDate" binding:"required"`
	PurchaseTime string  `json:"purchaseTime" binding:"required"`
	Items        []Item  `json:"items" binding:"required"`
	Total        float64 `json:"total" binding:"required"`
	Points       int     `json:"points,omitempty"`
}

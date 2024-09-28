package model

type Volume struct {
	Category      int     `json:"category" binding:"required,numeric"`
	Amount        int     `json:"amount" binding:"required,numeric"`
	UnitaryWeight int     `json:"unitary_weight" binding:"required,numeric"`
	Price         float64 `json:"price" binding:"required,numeric"`
	Sku           string  `json:"sku" binding:"required"`
	Height        float64 `json:"height" binding:"required,numeric"`
	Width         float64 `json:"width" binding:"required,numeric"`
	Length        float64 `json:"length" binding:"required,numeric"`
}

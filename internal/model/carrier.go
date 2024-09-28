package model

type Carrier struct {
	ID       int     `json:"id_carrier"`
	Name     string  `json:"name" binding:"required"`
	Service  string  `json:"service" binding:"required"`
	Deadline int     `json:"deadline" binding:"required,numeric"`
	Price    float64 `json:"price" binding:"required,numeric"`
}

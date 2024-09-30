package model

type Metrics struct {
	Quantity     int       `json:"quantity"`
	TotalPrice   float64   `json:"total_price"`
	AveragePrice float64   `json:"average_price"`
	Cheaper      Carrier   `json:"cheaper"`
	Expensive    Carrier   `json:"expensive"`
	Carriers     []Carrier `json:"carriers"`
}

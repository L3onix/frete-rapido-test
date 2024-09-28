package resbody

type QuoteSimulateV3ResBody struct {
	Dispatchers []Dispatcher `json:"dispatchers"`
}

type Dispatcher struct {
	ID         string  `json:"id"`
	Offers     []Offer `json:"offers"`
	TotalPrice float64 `json:"total_price"`
}

type Offer struct {
	ID           int          `json:"offer"`
	Carrier      Carrier      `json:"carrier"`
	Service      string       `json:"service"`
	DeliveryTime DeliveryTime `json:"delivery_time"`
	FinalPrice   float64      `json:"final_price"`
}

type Carrier struct {
	Name string `json:"name"`
}

type DeliveryTime struct {
	Days int `json:"days"`
}

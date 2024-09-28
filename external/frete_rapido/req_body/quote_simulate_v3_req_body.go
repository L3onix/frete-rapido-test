package reqbody

import (
	"fmt"
	"strconv"

	"github.com/L3onix/frete-rapido-test/internal/model"
)

type QuoteSimulateV3ReqBody struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	Channel        string       `json:"channel"`
	Filter         int          `json:"filter"`
	Limit          int          `json:"limit"`
	Identification string       `json:"identification"`
	Reverse        bool         `json:"reverse"`
	SimulationType []int        `json:"simulation_type"`
	Returns        Returns      `json:"returns"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Country          string `json:"country"`
	Zipcode          int    `json:"zipcode"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int      `json:"zipcode"`
	TotalPrice       float64  `json:"total_price"`
	Volumes          []Volume `json:"volumes"`
}

type Volume struct {
	Category      string  `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight int     `json:"unitary_weight"`
	Price         float64 `json:"price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type Returns struct {
	Composition  bool `json:"composition"`
	Volumes      bool `json:"volumes"`
	AppliedRules bool `json:"applied_rules"`
}

func LoadReqBody(quote model.Quote) (QuoteSimulateV3ReqBody, error) {
	zipcode, err := strconv.Atoi(quote.Recipient.Address.Zipcode)
	if err != nil {
		fmt.Println(err)
		return QuoteSimulateV3ReqBody{}, err
	}

	shipper := Shipper{
		RegisteredNumber: "25438296000158",
		Token:            "1d52a9b6b78cf07b08586152459a5c90",
		PlatformCode:     "5AKVkHqCn",
	}

	recipient := Recipient{
		Type:             0,
		RegisteredNumber: "",
		StateInscription: "",
		Country:          "BRA",
		Zipcode:          zipcode,
	}

	var volumes []Volume
	var total_price float64
	for _, quoteVolume := range quote.Volumes {
		category := strconv.Itoa(quoteVolume.Category)

		volume := Volume{
			Category:      category,
			Amount:        quoteVolume.Amount,
			UnitaryPrice:  quoteVolume.Price,
			UnitaryWeight: quoteVolume.UnitaryWeight,
			Price:         quoteVolume.Price,
			Sku:           quoteVolume.Sku,
			Height:        quoteVolume.Height,
			Width:         quoteVolume.Width,
			Length:        quoteVolume.Length,
		}

		volumes = append(volumes, volume)
		total_price += quoteVolume.Price
	}

	dispatchers := []Dispatcher{
		{
			RegisteredNumber: "25438296000158",
			Zipcode:          29161376,
			TotalPrice:       total_price,
			Volumes:          volumes,
		},
	}

	returns := Returns{
		Composition:  false,
		Volumes:      false,
		AppliedRules: false,
	}

	return QuoteSimulateV3ReqBody{
		Shipper:        shipper,
		Recipient:      recipient,
		Dispatchers:    dispatchers,
		Channel:        "",
		Filter:         0,
		Limit:          0,
		Identification: "",
		Reverse:        false,
		SimulationType: []int{0},
		Returns:        returns,
	}, nil
}

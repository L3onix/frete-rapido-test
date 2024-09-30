package usecase

import (
	"fmt"
	"strconv"

	"github.com/L3onix/frete-rapido-test/internal/model"
)

type MetricUsecase struct {
	carrierUseCase CarrierUseCase
}

func NewMetricUseCase(carrierUseCase CarrierUseCase) MetricUsecase {
	return MetricUsecase{
		carrierUseCase: carrierUseCase,
	}
}

func (mu *MetricUsecase) GetMetrics(lastQuotes string, dispatcherID string) (model.Metrics, error) {
	carriers, err := mu.carrierUseCase.GetLastCarriers(lastQuotes, dispatcherID)
	if err != nil {
		return model.Metrics{}, err
	}

	var metrics model.Metrics
	slice := 0
	if lastQuotes != "" {
		slice, err = strconv.Atoi(lastQuotes)
		if err != nil {
			fmt.Printf("Erro ao converter lastQuotes: %s", err)
			return model.Metrics{}, err
		}
	}

	metrics = mu.loadMetrics(*carriers, slice)
	return metrics, err
}

func (mu *MetricUsecase) loadMetrics(carriers []model.Carrier, slice int) model.Metrics {

	var totalPrice float64
	cheaper := carriers[0]
	expansive := carriers[0]
	var lastCarriers []model.Carrier

	for index, carrier := range carriers {
		totalPrice += carrier.Price
		if carrier.Price < cheaper.Price {
			cheaper = carrier
		}
		if carrier.Price > expansive.Price {
			expansive = carrier
		}
		if slice > 0 && slice > index {
			lastCarriers = append(lastCarriers, carrier)
		}
	}
	if slice == 0 {
		lastCarriers = carriers
	}

	averagePrice := totalPrice / float64(len(carriers))

	return model.Metrics{
		Quantity:     len(carriers),
		TotalPrice:   totalPrice,
		AveragePrice: averagePrice,
		Cheaper:      cheaper,
		Expensive:    expansive,
		Carriers:     lastCarriers,
	}
}

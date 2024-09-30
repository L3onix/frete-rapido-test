package usecase

import "github.com/L3onix/frete-rapido-test/internal/model"

type MetricUsecase struct {
	carrierUseCase CarrierUseCase
}

func NewMetricUseCase(carrierUseCase CarrierUseCase) MetricUsecase {
	return MetricUsecase{
		carrierUseCase: carrierUseCase,
	}
}

func (mu *MetricUsecase) GetMetrics(lastQuotes string, dispatcherID string) (*[]model.Carrier, error) {
	return mu.carrierUseCase.GetLastCarriers(lastQuotes, dispatcherID)
}

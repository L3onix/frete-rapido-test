package usecase

import (
	"github.com/L3onix/frete-rapido-test/internal/model"
	"github.com/L3onix/frete-rapido-test/internal/repository"
)

type CarrierUseCase struct {
	repository repository.CarrierRepository
}

func NewCarrierUseCase(repository repository.CarrierRepository) CarrierUseCase {
	return CarrierUseCase{
		repository: repository,
	}
}

func (cu *CarrierUseCase) GetCarriers() ([]model.Carrier, error) {
	return cu.repository.GetCarriers()
}

func (cu *CarrierUseCase) CreateCarrier(carrier model.Carrier) (model.Carrier, error) {
	id, err := cu.repository.CreateCarrier(carrier)
	if err != nil {
		return model.Carrier{}, err
	}

	carrier.ID = id
	return carrier, err
}

func (cu *CarrierUseCase) GetCarrierById(id int) (*model.Carrier, error) {
	carrier, err := cu.repository.GetCarrierById(id)
	if err != nil {
		return nil, err
	}

	return carrier, nil
}

func (cu *CarrierUseCase) GetLastCarriers(lastCarriers string) (*[]model.Carrier, error) {
	return cu.repository.GetLastCarriers(lastCarriers)
}

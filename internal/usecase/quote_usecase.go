package usecase

import (
	resbody "github.com/L3onix/frete-rapido-test/external/frete_rapido/res_body"
	"github.com/L3onix/frete-rapido-test/internal/model"
)

type QuoteUseCase struct {
	carrierUseCase CarrierUseCase
}

func NewQuoteUseCase(CarrierUseCase CarrierUseCase) QuoteUseCase {
	return QuoteUseCase{
		carrierUseCase: CarrierUseCase,
	}
}

func (qu *QuoteUseCase) StoreResult(simResult resbody.QuoteSimulateV3ResBody) {
	var carrier model.Carrier
	panic(carrier)
}

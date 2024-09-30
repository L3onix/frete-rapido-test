package usecase

import (
	"database/sql"

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

func (qu *QuoteUseCase) StoreResult(simResult resbody.QuoteSimulateV3ResBody) ([]model.Carrier, error) {
	carriers := []model.Carrier{}
	var err error
	for _, offer := range simResult.Dispatchers[0].Offers {
		carrier, err := qu.carrierUseCase.CreateCarrier(model.Carrier{
			Name:         offer.Carrier.Name,
			Service:      offer.Service,
			Deadline:     offer.DeliveryTime.Days,
			Price:        offer.FinalPrice,
			DispatcherID: sql.NullString{String: simResult.Dispatchers[0].ID, Valid: true},
		})
		if err != nil {
			carriers = []model.Carrier{}
			break
		}

		carriers = append(carriers, carrier)
	}

	return carriers, err
}

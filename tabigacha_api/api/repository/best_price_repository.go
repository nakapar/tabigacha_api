package repository

import (
	entity "tabigacha-api/api/entity/flight"
)

type IFetchFlightPriceRepository interface {
	FetchFlightPrice(entity *entity.FetchFlightPriceCondition) ([]entity.FlightEntity, error)
}

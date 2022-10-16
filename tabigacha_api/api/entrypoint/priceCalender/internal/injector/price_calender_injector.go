package injector

import (
	"tabigacha-api/api/config"
	"tabigacha-api/api/entrypoint/priceCalender/internal/service"
	"tabigacha-api/api/infra/awsdynamodb"
	"tabigacha-api/api/infra/awsdynamodb/flight_price"
)

type priceCalenderInjector struct {
	DynamodbClient awsdynamodb.IDynamodbClient
	TableName      string
}

const tableName = "flight_price_t"

func NewPriceCalenderInjector() (*priceCalenderInjector, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	dynamodbc, err := awsdynamodb.NewDynamodbClient(conf, tableName)
	if err != nil {
		return nil, err
	}

	return &priceCalenderInjector{
		DynamodbClient: dynamodbc,
		TableName:      tableName,
	}, nil
}

func (i *priceCalenderInjector) NewService() service.PriceCalenderService {
	repo := flight_price.NewFlightPriceRepositoryDynamo(i.DynamodbClient)
	return service.NewPriceCalenderService(repo)
}

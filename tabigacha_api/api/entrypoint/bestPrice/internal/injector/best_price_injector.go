package injector

import (
	"tabigacha-api/api/config"
	"tabigacha-api/api/entrypoint/bestPrice/internal/service"
	"tabigacha-api/api/infra/awsdynamodb"
	"tabigacha-api/api/infra/awsdynamodb/flight_price"
)

type bestPriceInjector struct {
	DynamodbClient awsdynamodb.IDynamodbClient
	TableName      string
}

const tableName = "flight_price_t"

func NewBestPriceInjector() (*bestPriceInjector, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	dynamodbc, err := awsdynamodb.NewDynamodbClient(conf, tableName)
	if err != nil {
		return nil, err
	}

	return &bestPriceInjector{
		DynamodbClient: dynamodbc,
		TableName: tableName,
	}, nil
}

func (i *bestPriceInjector) NewService() service.BestPriceService {
	repo := flight_price.NewFlightPriceRepositoryDynamo(i.DynamodbClient)
	return service.NewBestPriceService(repo)
}

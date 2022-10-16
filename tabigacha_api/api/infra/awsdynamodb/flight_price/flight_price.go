package flight_price

import (
	entity "tabigacha-api/api/entity/flight"
	"tabigacha-api/api/infra/awsdynamodb"
	"tabigacha-api/api/repository"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type flightPriceRepositoryDynamo struct {
	client awsdynamodb.IDynamodbClient
}

func NewFlightPriceRepositoryDynamo(client awsdynamodb.IDynamodbClient) repository.IFetchFlightPriceRepository {
	return &flightPriceRepositoryDynamo{client: client}
}

func (d *flightPriceRepositoryDynamo) FetchFlightPrice(cond *entity.FetchFlightPriceCondition) ([]entity.FlightEntity, error) {
	var flightEntity []entity.FlightEntity

	params := awsdynamodb.NewRequestData(cond.KeyConditionExpression, cond.ExpressionAttributeValues)

	output, err := d.client.QueryFetch(params)
	if err != nil {
		return nil, err
	}

	fe := entity.FlightEntity{}

	for _, item := range output.Items {
		err = attributevalue.UnmarshalMap(item, &fe)
		if err != nil {
			return nil, err
		}
		flightEntity = append(flightEntity, fe)
	}
	return flightEntity, nil
}

package service

import (
	"strconv"
	entity "tabigacha-api/api/entity/flight"
	"tabigacha-api/api/entrypoint/priceCalender/internal/form"
	"tabigacha-api/api/entrypoint/priceCalender/internal/response"
	"tabigacha-api/api/repository"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PriceCalenderService interface {
	FetchPriceCalender(form *form.PriceCalenderForm) (*response.PriceCalenderResponse, error)
}

type priceCalenderService struct {
	fbpr repository.IFetchFlightPriceRepository
}

func NewPriceCalenderService(fbpr repository.IFetchFlightPriceRepository) PriceCalenderService {
	return &priceCalenderService{fbpr: fbpr}
}

func (b *priceCalenderService) FetchPriceCalender(form *form.PriceCalenderForm) (*response.PriceCalenderResponse, error) {
	keyConditionExpression := `Origin = :origin And Destination = :destination`
	expressionAttributeValues := map[string]types.AttributeValue{
		":origin":      &types.AttributeValueMemberS{Value: form.Origin},
		":destination": &types.AttributeValueMemberS{Value: form.Destination},
	}

	cond := &entity.FetchFlightPriceCondition{
		KeyConditionExpression:    keyConditionExpression,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	resp, err := b.fbpr.FetchFlightPrice(cond)
	if err != nil {
		return nil, err
	}

	return toResponse(resp), nil
}

func toResponse(resp []entity.FlightEntity) *response.PriceCalenderResponse {
	priceCal := resp[0].PriceCalender
	datePrices := []response.DatePrice{}
	for date, price := range priceCal {
		intPrice, _ := strconv.Atoi(price)
		datePrices = append(datePrices, response.DatePrice{
			Date: date,
			Price: intPrice,
		})
	}
	return &response.PriceCalenderResponse{
		Dates: datePrices,
	}
}
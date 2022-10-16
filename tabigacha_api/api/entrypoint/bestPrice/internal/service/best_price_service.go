package service

import (
	"strconv"
	entity "tabigacha-api/api/entity/flight"
	"tabigacha-api/api/entrypoint/bestPrice/internal/form"
	"tabigacha-api/api/entrypoint/bestPrice/internal/response"
	"tabigacha-api/api/repository"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BestPriceService interface {
	FetchBestPrice(form *form.BestPriceForm) (*response.BestPriceResponse, error)
}

type bestPriceService struct {
	fbpr repository.IFetchFlightPriceRepository
}

func NewBestPriceService(fbpr repository.IFetchFlightPriceRepository) BestPriceService {
	return &bestPriceService{fbpr: fbpr}
}

func (b *bestPriceService) FetchBestPrice(form *form.BestPriceForm) (*response.BestPriceResponse, error) {
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

func toResponse(resp []entity.FlightEntity) *response.BestPriceResponse {
	flight := resp[0].BestPrice
	bestPrice, _ := strconv.Atoi(flight["best_price"])
	date := searchBestPrice(resp)
	return &response.BestPriceResponse{
		Price: bestPrice,
		Date:  date,
	}
}

func searchBestPrice(resp []entity.FlightEntity) []string {
	var bestPriceDate []string
	bestPrice, _ := strconv.Atoi(resp[0].BestPrice["best_price"])
	priceCal := resp[0].PriceCalender
	for date, price := range priceCal {
		intPrice, _ := strconv.Atoi(price)
		if intPrice == bestPrice {
			bestPriceDate = append(bestPriceDate, date)
		}
	}
	return bestPriceDate
}

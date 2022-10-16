package entity

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type FetchFlightPriceCondition struct {
	KeyConditionExpression    string
	ExpressionAttributeValues map[string]types.AttributeValue
}

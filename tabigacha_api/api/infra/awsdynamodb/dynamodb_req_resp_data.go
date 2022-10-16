package awsdynamodb

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type queryFetchRequest struct {
	KeyConditionExpression    string
	ExpressionAttributeValues map[string]types.AttributeValue
}

func NewRequestData(keyConditionExpression string, expressionAttributeValues map[string]types.AttributeValue) *queryFetchRequest {
	return &queryFetchRequest{
		KeyConditionExpression: keyConditionExpression,
		ExpressionAttributeValues: expressionAttributeValues,
	}
}

type QueryFetchOutput struct {
	Count int32
	Items []map[string]types.AttributeValue
}

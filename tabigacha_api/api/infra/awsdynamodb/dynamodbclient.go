package awsdynamodb

import (
	"context"
	"tabigacha-api/api/config"

	awsConf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"golang.org/x/xerrors"
)

type IDynamodbClient interface {
	QueryFetch(*queryFetchRequest) (dynamodb.QueryOutput, error)
}

type DynamodbClient struct {
	Client    *dynamodb.Client
	TableName string
}

func NewDynamodbClient(config *config.Config, tableName string) (*DynamodbClient, error) {
	awsConf, err := awsConf.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	var client *dynamodb.Client

	if config.IsLocal() {
		endpointResolver := dynamodb.EndpointResolverFromURL(config.AwsDynamodb.Endpoint)
		client = dynamodb.NewFromConfig(awsConf,
			dynamodb.WithEndpointResolver(endpointResolver),
		)
	} else {
		client = dynamodb.NewFromConfig(awsConf)
	}

	return &DynamodbClient{Client: client, TableName: tableName}, nil
}

func (c *DynamodbClient) QueryFetch(request *queryFetchRequest) (dynamodb.QueryOutput, error) {
	result, err := c.Client.Query(
		context.TODO(),
		&dynamodb.QueryInput{
			TableName:                 &c.TableName,
			KeyConditionExpression:    &request.KeyConditionExpression,
			ExpressionAttributeValues: request.ExpressionAttributeValues,
		},
	)
	if err != nil {
		return dynamodb.QueryOutput{}, xerrors.Errorf(err.Error())
	}
	return *result, err
}

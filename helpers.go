package main

// Please check https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
// for best practices setting up your AWS credentials

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ConnectionItem struct {
	ConnectionID string `json:"connectionID"`
}

const (
	AccessKeyID        = "YOUR_ACCESS_KEY"
	SecretAccessKey    = "YOUR_SECRET_KEY"
	APIGatewayEndpoint = "YOUR_API_GATEWAY_ENDPOINT"
	Region             = "YOUR_REGION"
)

func NewDynamoDBSession() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	})
	return dynamodb.New(sess)
}

func NewAPIGatewaySession() *apigatewaymanagementapi.ApiGatewayManagementApi {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		Endpoint:    aws.String(APIGatewayEndpoint),
	})
	return apigatewaymanagementapi.New(sess)
}

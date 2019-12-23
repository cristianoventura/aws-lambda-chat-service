package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	lambda.Start(HandleMessage)
}

// RequestPayload represents the request body sent by the Socket
type RequestPayload struct {
	Message string `json:"message"`
}

type MessageData struct {
	Message      string `json:"message"`
	ConnectionID string `json:"connectionId"`
}

func HandleMessage(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the request
	var requestPayload RequestPayload
	json.Unmarshal([]byte(request.Body), &requestPayload)

	dynamodbSession := NewDynamoDBSession()

	// Read the chat-connections table
	input := &dynamodb.ScanInput{
		TableName: aws.String("chat-connections"),
	}
	scan, _ := dynamodbSession.Scan(input)

	// Parse the table data in the output variable
	var output []ConnectionItem
	dynamodbattribute.UnmarshalListOfMaps(scan.Items, &output)

	apigatewaySession := NewAPIGatewaySession()

	// Encode the message data with Message and Connection ID
	messageData := &MessageData{
		Message:      requestPayload.Message,
		ConnectionID: request.RequestContext.ConnectionID,
	}
	jsonData, _ := json.Marshal(messageData)

	// Send the message for each connection ID
	for _, item := range output {
		connectionInput := &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(item.ConnectionID),
			Data:         jsonData,
		}
		_, err := apigatewaySession.PostToConnection(connectionInput)
		if err != nil {
			fmt.Println(err)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}

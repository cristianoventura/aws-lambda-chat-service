# aws-lambda-chat-service
Source code for the Create a Real Time Chat Service with Go, API Gateway and Lambda Tutorial

You can find the complete tutorial written by me here: https://medium.com/@cristiano.ventura/create-a-real-time-chat-service-with-go-api-gateway-and-lambda-959f8c9752e6

## Building the Functions

There are two functions:
+ chat-message.go
+ chat-connect.go

AWS Lambda expects two runtime version:
GOOS=linux and GOARCH=amd64

So you can build it like this:

`export GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && go build -o ./main chat-message.go helpers.go`

For this example, the `helpers.go` file is a dependency for both functions.

Zip them and upload to you Lambda Function

## Using the `client.html` file

Replace the YOUR_WEBSOCKET_URL string with your API Gateway Connection URL.

```javascript
const socket = new WebSocket('YOUR_WEBSOCKET_URL');
```

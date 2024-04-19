package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"strings"
)

type RequestBody struct {
	Test string `json:"test"`
}

func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// Normalize path
	path := strings.Trim(request.RawPath, "/")

	// Default response for "/"
	if path == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Hello from Root!",
		}, nil
	}

	switch path {
	case "reflect":
		testValue := request.QueryStringParameters["test"]
		if testValue == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest, // 400
				Body:       "Query parameter 'test' is missing",
			}, nil
		}
		responseMessage := "Received 'test' parameter: " + testValue
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       responseMessage,
		}, nil
	case "post":
		var body RequestBody
		err := json.Unmarshal([]byte(request.Body), &body)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Invalid JSON in request body",
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf("POST request received with 'test' value: %s", body.Test),
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}

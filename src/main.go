package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	nf "github.com/habibiefaried/golang-web-lambda/library/networkfirewall"
	"net/http"
	"strings"
)

type RequestBody struct {
	Domain string `json:"domain"`
}

func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// Normalize path
	path := strings.Trim(request.RawPath, "/")

	// Default response for "/"
	if path == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Hello World!",
		}, nil
	}

	switch path {
	case "is-whitelisted":
		domain := request.QueryStringParameters["domain"]
		if domain == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Query parameter 'domain' is missing",
			}, nil
		}

		isWhitelisted, err := nf.IsDomainWhitelisted("test", domain)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       err.Error(),
			}, nil
		} else {
			responseMessage := "yes"
			if !isWhitelisted {
				responseMessage = "no"
			}
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       responseMessage,
			}, nil
		}
	case "whitelist":
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
			Body:       fmt.Sprintf("POST request received with 'domain' value: %s", body.Domain),
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

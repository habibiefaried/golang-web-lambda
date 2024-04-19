package main

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
    "strings"
)

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
    case "test":
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       "Hello from Test!",
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

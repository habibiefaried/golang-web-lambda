package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	nf "github.com/habibiefaried/golang-web-lambda/library/networkfirewall"
	"net/http"
	"os"
	"strings"
	"log"
)

func handleRequest(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	log.Printf("Received request: %+v", request)

	// Normalize path
    NetworkFirewallRuleGroupName := os.Getenv("RULEGROUPNAME")
    path := strings.Trim(request.Path, "/ ")
    
    // Default response for "/"
    if path == "" {
        return events.ALBTargetGroupResponse{
            StatusCode: http.StatusOK,
            Body:       "Hello World!",
            Headers:    map[string]string{"Content-Type": "text/plain"},
        }, nil
    }

    switch path {
    case "is-whitelisted":
        domain := request.QueryStringParameters["domain"]
        if domain == "" {
            return events.ALBTargetGroupResponse{
                StatusCode: http.StatusBadRequest,
                Body:       "Query parameter 'domain' is missing",
                Headers:    map[string]string{"Content-Type": "text/plain"},
            }, nil
        }

        isWhitelisted, err := nf.IsDomainWhitelisted(NetworkFirewallRuleGroupName, domain)
        if err != nil {
            return events.ALBTargetGroupResponse{
                StatusCode: http.StatusInternalServerError,
                Body:       err.Error(),
                Headers:    map[string]string{"Content-Type": "text/plain"},
            }, nil
        } else {
            responseMessage := "yes"
            if !isWhitelisted {
                responseMessage = "no"
            }
            return events.ALBTargetGroupResponse{
                StatusCode: http.StatusOK,
                Body:       responseMessage,
                Headers:    map[string]string{"Content-Type": "text/plain"},
            }, nil
        }
    case "whitelist":
        var body RequestBody
        err := json.Unmarshal([]byte(request.Body), &body)
        if err != nil {
            return events.ALBTargetGroupResponse{
                StatusCode: http.StatusBadRequest,
                Body:       "Invalid JSON in request body",
                Headers:    map[string]string{"Content-Type": "text/plain"},
            }, nil
        }

        token, err := nf.AddRule(NetworkFirewallRuleGroupName, body.Domain)
        if err != nil {
            return events.ALBTargetGroupResponse{
                StatusCode: http.StatusInternalServerError,
                Body:       err.Error(),
                Headers:    map[string]string{"Content-Type": "text/plain"},
            }, nil
        }

        return events.ALBTargetGroupResponse{
            StatusCode: http.StatusOK,
            Body:       fmt.Sprintf("Added domain %v to be whitelisted with token ref %v", body.Domain, *token),
            Headers:    map[string]string{"Content-Type": "text/plain"},
        }, nil
    default:
        return events.ALBTargetGroupResponse{
            StatusCode: http.StatusNotFound,
            Body:       "Not Found",
            Headers:    map[string]string{"Content-Type": "text/plain"},
        }, nil
    }
}

func main() {
	lambda.Start(handleRequest)
}

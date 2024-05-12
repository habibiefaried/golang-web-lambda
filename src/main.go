package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	apiweb "github.com/habibiefaried/golang-web-lambda/api/web"
)

// Handler is the Lambda function handler
func Handler(ctx context.Context, event json.RawMessage) (interface{}, error) {
	// Attempt to unmarshal the event as an ALB event
	var albEvent events.ALBTargetGroupRequest
	if err := json.Unmarshal(event, &albEvent); err == nil {
		if albEvent.HTTPMethod != "" {
			// Handle ALB event
			return apiweb.HandleRequestWeb(ctx, albEvent)
		}
	}

	// Attempt to unmarshal the event as an EventBridge event
	var ebEvent events.CloudWatchEvent
	if err := json.Unmarshal(event, &ebEvent); err == nil {
		if ebEvent.Source != "" {
			// Handle EventBridge event
			return handleEventBridgeEvent(ebEvent)
		}
	}

	// Fallback if the event does not match expected types
	return nil, fmt.Errorf("unrecognized event type")
}

func handleEventBridgeEvent(event events.CloudWatchEvent) (string, error) {
	// Define processing logic for EventBridge events here
	return "Hello from EventBridge", nil
}

func main() {
	lambda.Start(Handler)
}

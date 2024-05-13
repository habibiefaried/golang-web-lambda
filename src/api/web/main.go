package apiweb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	nf "github.com/habibiefaried/golang-web-lambda/library/networkfirewall"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleRequestWeb(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
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
	case "whitelist":
		var body WhitelistRequest
		err := json.Unmarshal([]byte(request.Body), &body)
		if err != nil {
			return events.ALBTargetGroupResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Invalid JSON in request body",
				Headers:    map[string]string{"Content-Type": "text/plain"},
			}, nil
		}
		log.Printf("Log %+v", body)
		err = nf.ManageRule(
			NetworkFirewallRuleGroupName,
			nf.RequestBody{
				ID:  body.ID,
				URL: body.OldURL,
			},
			nf.RequestBody{
				ID:  body.ID,
				URL: body.NewURL,
			},
		)
		if err != nil {
			return events.ALBTargetGroupResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       err.Error(),
				Headers:    map[string]string{"Content-Type": "text/plain"},
			}, nil
		}

		return events.ALBTargetGroupResponse{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf("Rule is updated!"),
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

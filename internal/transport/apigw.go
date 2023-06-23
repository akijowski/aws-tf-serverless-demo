package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"cdr.dev/slog"
	"github.com/aws/aws-lambda-go/events"
)

// APIGatewayTransport facilitates sending requests from Lambda to API Gateway
type APIGatewayTransport struct {
	logger slog.Logger
}

// NewAPIGateway creates a transport to be used with API Gateway Lambda events
func NewAPIGateway(logger slog.Logger) *APIGatewayTransport {
	return &APIGatewayTransport{logger}
}

// apiError would normally be exported as a module for clients or something other than in here
type apiError struct {
	Message string `json:"message"`
}

var jsonHeaders = map[string]string{"Content-Type": "application/json"}

// Send marshals the information in to a response for API Gateway.
// Any errors will route to SendError with a status code of 500.
func (t *APIGatewayTransport) Send(ctx context.Context, statusCode int, body any) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return t.SendError(ctx, http.StatusInternalServerError, err)
	}
	t.logger.Info(ctx, "request complete", slog.F("status_code", statusCode))
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    jsonHeaders,
		Body:       string(b),
	}, nil
}

// SendError marshals the information in to a response for API Gateway
func (t *APIGatewayTransport) SendError(ctx context.Context, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	t.logger.Warn(ctx, "request completed with error", slog.F("status_code", statusCode), slog.Error(err))
	resp := apiError{Message: err.Error()}
	b, err := json.Marshal(&resp)
	if err != nil {
		t.logger.Error(ctx, "error marshing response body", slog.Error(err))
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    jsonHeaders,
		Body:       string(b),
	}, nil
}

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
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(b),
	}, nil
}

// SendError marshals the information in to a response for API Gateway
func (t *APIGatewayTransport) SendError(ctx context.Context, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	t.logger.Warn(ctx, "request completed with error", slog.F("status_code", statusCode), slog.F("error", err))
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       err.Error(),
	}, nil
}

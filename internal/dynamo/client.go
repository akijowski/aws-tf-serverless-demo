package dynamo

import (
	"context"
	"os"

	"cdr.dev/slog"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// NewClient returns a dynamodb.Client.
// If DYNAMODB_ENDPOINT is set in the environment, the endpoint will be updated.  This is useful for tests.
func NewClient(ctx context.Context, logger slog.Logger) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	opts := []func(*dynamodb.Options){}

	endpointURI := os.Getenv("DYNAMODB_ENDPOINT")
	if endpointURI != "" {
		logger.Warn(ctx, "overriding dynamodb endpoint", slog.F("endpoint", endpointURI))
		opts = append(opts, dynamodb.WithEndpointResolver(dynamodb.EndpointResolverFromURL(endpointURI)))
	}

	return dynamodb.NewFromConfig(cfg, opts...), nil
}

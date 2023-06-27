package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"cdr.dev/slog"
	"github.com/akijowski/aws-tf-serverless-demo/internal/transport"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type CreateKeyStore interface {
	CreateIfNotExists(context.Context, *types.KeyValueEntry) error
}

type CreateKeyEntryTask struct {
	logger slog.Logger
	store  CreateKeyStore
}

func NewCreateKeyEntry(logger slog.Logger, store CreateKeyStore) *CreateKeyEntryTask {
	return &CreateKeyEntryTask{
		logger: logger,
		store:  store,
	}
}

func (t *CreateKeyEntryTask) HandleCreateKeyAPIEvent(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = withReqContext(ctx, req.RequestContext)
	ctx = withReqInfo(ctx, req)
	t.logger.Info(ctx, "handling request")
	trans := transport.NewAPIGateway(t.logger)
	var entry *types.KeyValueEntry
	if err := json.Unmarshal([]byte(req.Body), &entry); err != nil {
		return trans.SendError(ctx, http.StatusBadRequest, err)
	}
	if entry.Key == "" {
		return trans.SendError(ctx, http.StatusBadRequest, errors.New("key is required"))
	}
	if err := t.store.CreateIfNotExists(ctx, entry); err != nil {
		return trans.SendError(ctx, http.StatusInternalServerError, err)
	}

	return trans.Send(ctx, http.StatusCreated, entry)
}

func withReqContext(ctx context.Context, reqCtx events.APIGatewayProxyRequestContext) context.Context {
	updatedCtx := slog.With(ctx, slog.F("api_request_id", reqCtx.RequestID))
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		updatedCtx = slog.With(updatedCtx, slog.F("lambda_request_id", lc.AwsRequestID))
	}
	return updatedCtx
}

func withReqInfo(ctx context.Context, req events.APIGatewayProxyRequest) context.Context {
	updatedCtx := slog.With(ctx, slog.F("path", req.Path), slog.F("method", req.HTTPMethod))
	if len(req.PathParameters) > 0 {
		updatedCtx = slog.With(ctx, slog.F("path_params", req.PathParameters))
	}

	return updatedCtx
}

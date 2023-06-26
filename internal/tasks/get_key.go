package tasks

import (
	"context"
	"errors"
	"net/http"

	"cdr.dev/slog"
	"github.com/akijowski/aws-tf-serverless-demo/internal/transport"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-lambda-go/events"
)

type GetKeyStore interface {
	GetEntryByKey(context.Context, string) (*types.KeyValueEntry, error)
}

type GetKeyEntryTask struct {
	logger slog.Logger
	store  GetKeyStore
}

func NewGetKeyEntry(logger slog.Logger, store GetKeyStore) *GetKeyEntryTask {
	return &GetKeyEntryTask{
		logger: logger,
		store:  store,
	}
}

func (t *GetKeyEntryTask) HandleGetKeyAPIRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = withReqContext(ctx, req.RequestContext)
	ctx = withReqInfo(ctx, req)

	t.logger.Info(ctx, "handling request")
	trans := transport.NewAPIGateway(t.logger)

	keyId, ok := req.PathParameters["keyID"]
	if !ok {
		return trans.SendError(ctx, http.StatusBadRequest, errors.New("missing required path param: keyID"))
	}
	found, err := t.store.GetEntryByKey(ctx, keyId)
	if err != nil {
		return trans.SendError(ctx, http.StatusInternalServerError, err)
	}
	if found == nil {
		return trans.SendError(ctx, http.StatusNotFound, errors.New("entry not found"))
	}

	return trans.Send(ctx, http.StatusOK, found)
}

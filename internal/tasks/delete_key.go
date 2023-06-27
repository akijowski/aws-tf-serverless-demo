package tasks

import (
	"context"
	"errors"
	"net/http"

	"cdr.dev/slog"
	"github.com/akijowski/aws-tf-serverless-demo/internal/transport"
	"github.com/aws/aws-lambda-go/events"
)

type DeleteKeyStore interface {
	DeleteEntryByKey(context.Context, string) error
}

type DeleteKeyEntryTask struct {
	logger slog.Logger
	store  DeleteKeyStore
}

func NewDeleteKeyEntry(logger slog.Logger, store DeleteKeyStore) *DeleteKeyEntryTask {
	return &DeleteKeyEntryTask{
		logger: logger,
		store:  store,
	}
}

func (t *DeleteKeyEntryTask) HandleDeleteKeyAPIEvent(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = withReqContext(ctx, req.RequestContext)
	ctx = withReqInfo(ctx, req)

	t.logger.Info(ctx, "handling request")
	trans := transport.NewAPIGateway(t.logger)

	key, ok := req.PathParameters["keyID"]
	if !ok {
		return trans.SendError(ctx, http.StatusBadRequest, errors.New("missing required path param: keyID"))
	}

	err := t.store.DeleteEntryByKey(ctx, key)
	if err != nil {
		return trans.SendError(ctx, http.StatusInternalServerError, err)
	}

	return trans.Send(ctx, http.StatusNoContent, nil)
}

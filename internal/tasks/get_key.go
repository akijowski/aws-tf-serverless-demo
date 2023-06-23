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
	GetItem(context.Context, string) (*types.KeyValueEntry, error)
}

type GetKeyEntryTask struct {
	logger slog.Logger
	store  GetKeyStore
}

func (t *GetKeyEntryTask) HandleGetKeyAPIRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = withReqContext(ctx, req.RequestContext)

	t.logger.Info(ctx, "handling request")
	trans := transport.NewAPIGateway(t.logger)

	keyId, ok := req.PathParameters["keyID"]
	if !ok {
		trans.SendError(ctx, http.StatusBadRequest, errors.New("missing required path param: keyID"))
	}
	found, err := t.store.GetItem(ctx, keyId)
	if err != nil {
		trans.SendError(ctx, http.StatusInternalServerError, err)
	}

	return trans.Send(ctx, http.StatusOK, found)
}
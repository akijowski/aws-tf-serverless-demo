package tasks

import (
	"context"
	"net/http"
	"testing"

	"cdr.dev/slog/sloggers/slogtest"
	"github.com/akijowski/aws-tf-serverless-demo/internal/store"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-lambda-go/events"
)

type mockCreateKeyStore func(context.Context, store.PutItemAPI, *types.KeyValueEntry) error

func (m mockCreateKeyStore) CreateIfNotExists(ctx context.Context, c store.PutItemAPI, e *types.KeyValueEntry) error {
	return m(ctx, c, e)
}

func TestHandleCreateKeyAPIEvent(t *testing.T) {

	testLogger := slogtest.Make(t, nil)

	cases := map[string]struct {
		given events.APIGatewayProxyRequest
		store func(*testing.T) CreateKeyStore
		want  events.APIGatewayProxyResponse
	}{
		"valid request calls store and returns": {
			given: events.APIGatewayProxyRequest{
				Body: `{"key": "abc", "value": "123"}`,
			},
			store: func(t *testing.T) CreateKeyStore {
				return mockCreateKeyStore(func(ctx context.Context, pia store.PutItemAPI, kve *types.KeyValueEntry) error {
					t.Helper()

					return nil
				})
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusCreated,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"key":"abc","value":"123"}`,
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			task := &CreateKeyEntryTask{
				logger: testLogger,
				store:  tt.store(t),
			}

			actual, err := task.HandleCreateKeyAPIEvent(context.Background(), tt.given)

			if err != nil {
				t.Fatal(err)
			}

			if actual.StatusCode != tt.want.StatusCode {
				t.Errorf("%d != %d\n", actual.StatusCode, tt.want.StatusCode)
			}

			if actual.Body != tt.want.Body {
				t.Errorf("%s != %s\n", actual.Body, tt.want.Body)
			}
		})
	}
}

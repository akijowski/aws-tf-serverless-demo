package tasks

import (
	"context"
	"net/http"
	"testing"

	"cdr.dev/slog/sloggers/slogtest"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-lambda-go/events"
)

type mockGetKeyStore func(context.Context, string) (*types.KeyValueEntry, error)

func (m mockGetKeyStore) GetItem(ctx context.Context, key string) (*types.KeyValueEntry, error) {
	return m(ctx, key)
}

func TestHandleGetKeyAPIRequest(t *testing.T) {

	testLogger := slogtest.Make(t, nil)

	cases := map[string]struct {
		given events.APIGatewayProxyRequest
		store func(*testing.T) GetKeyStore
		want  events.APIGatewayProxyResponse
	}{
		"valid request calls store and returns": {
			given: events.APIGatewayProxyRequest{
				PathParameters: map[string]string{"keyID": "abc"},
			},
			store: func(t *testing.T) GetKeyStore {
				return mockGetKeyStore(func(ctx context.Context, s string) (*types.KeyValueEntry, error) {
					t.Helper()

					return &types.KeyValueEntry{
						Key:   "abc",
						Value: "123",
					}, nil
				})
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"key":"abc","value":"123"}`,
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			task := &GetKeyEntryTask{
				logger: testLogger,
				store:  tt.store(t),
			}

			actual, err := task.HandleGetKeyAPIRequest(context.Background(), tt.given)

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

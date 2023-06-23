package store

import (
	"context"
	"testing"

	"cdr.dev/slog/sloggers/slogtest"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockPutItemAPI func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)

func (m mockPutItemAPI) PutItem(ctx context.Context, in *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m(ctx, in)
}

func TestCreateIfNotExists(t *testing.T) {

	testLogger := slogtest.Make(t, nil)

	cases := map[string]struct {
		given *types.KeyValueEntry
		db    func(*testing.T) PutItemAPI
		want  error
	}{
		"saves input to database": {
			given: &types.KeyValueEntry{
				Key:   "abc",
				Value: "123",
			},
			db: func(t *testing.T) PutItemAPI {
				return mockPutItemAPI(func(ctx context.Context, pii *dynamodb.PutItemInput, f ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
					t.Helper()

					return &dynamodb.PutItemOutput{}, nil
				})
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			store := &KeyValueStore{
				logger:    testLogger,
				tableName: "foo",
			}
			createStore := store.CreateStoreWith(tt.db(t))

			actual := createStore.CreateIfNotExists(context.Background(), tt.given)

			if actual != tt.want {
				t.Error("incorrect response")
				t.Logf("actual: %v\n", actual)
				t.Logf("wanted: %v\n", tt.want)
			}
		})
	}
}

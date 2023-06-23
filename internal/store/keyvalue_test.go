package store

import (
	"context"
	"testing"
	"time"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogtest"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockPutItemAPI func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)

func (m mockPutItemAPI) PutItem(ctx context.Context, in *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m(ctx, in)
}

type mockGetItemAPI func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)

func (m mockGetItemAPI) GetItem(ctx context.Context, in *dynamodb.GetItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
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

func TestGetEntryByKey(t *testing.T) {
	testLogger := slogtest.Make(t, nil).Leveled(slog.LevelDebug)

	cases := map[string]struct {
		given string
		db    func(*testing.T) GetItemAPI
		want  *types.KeyValueEntry
	}{
		"gets record from database": {
			given: "foo",
			db: func(t *testing.T) GetItemAPI {
				return mockGetItemAPI(func(ctx context.Context, gii *dynamodb.GetItemInput, f ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
					t.Helper()

					record := &keyValueDynamoRecord{
						PrimaryKey: "KEY#abc",
						Key:        "abc",
						Value:      "123",
						CreatedAt:  time.Now(),
					}
					item, err := attributevalue.MarshalMap(&record)
					if err != nil {
						t.Fatal(err)
					}

					return &dynamodb.GetItemOutput{Item: item}, nil
				})
			},
			want: &types.KeyValueEntry{
				Key:   "abc",
				Value: "123",
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			store := &KeyValueStore{
				logger:    testLogger,
				tableName: "foo",
			}
			getStore := store.GetStoreWith(tt.db(t))

			actual, err := getStore.GetEntryByKey(context.Background(), tt.given)
			if err != nil {
				t.Error(err)
			}
			if actual.Key != tt.want.Key {
				t.Errorf("incorrect key: %s\n", actual.Key)
			}

			if actual.Value != tt.want.Value {
				t.Errorf("incorrect value: %s\n", actual.Value)
			}
		})
	}
}

package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cdr.dev/slog"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PutItemAPI interface {
	PutItem(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type GetItemAPI interface {
	GetItem(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type DeleteItemAPI interface {
	DeleteItem(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

type keyValueDynamoRecord struct {
	PrimaryKey string    `json:"pk" dynamodbav:"pk"`
	Key        string    `json:"key" dynamodbav:"key"`
	Value      string    `json:"value" dynamodbav:"value"`
	CreatedAt  time.Time `json:"created_at" dynamodbav:"created_at"`
}

// KeyValueStore is a factory-like struct that provides functions to generate functional stores based on relevant CRUD operations.
type KeyValueStore struct {
	logger    slog.Logger
	tableName string
}

// With returns a factory for KeyValueStore.  Functional stores can then be created from this factory.
func With(logger slog.Logger, tableName string) *KeyValueStore {
	return &KeyValueStore{
		logger,
		tableName,
	}
}

type createKeyStore func(context.Context, *types.KeyValueEntry) error

// CreateIfNotExists conditionally saves a new types.KeyValueEntry if the item is not present based on the Key value.
func (s createKeyStore) CreateIfNotExists(ctx context.Context, entry *types.KeyValueEntry) error {
	return s(ctx, entry)
}

type getKeyStore func(context.Context, string) (*types.KeyValueEntry, error)

// GetEntryByKey retrieves an entry for a given key.  If no entry is found, returns nil.
func (s getKeyStore) GetEntryByKey(ctx context.Context, key string) (*types.KeyValueEntry, error) {
	return s(ctx, key)
}

type deleteKeyStore func(context.Context, string) error

// DeleteEntryByKey removes a record from the store for the given primary key.
// If the record does not exist, no error is returned.
func (s deleteKeyStore) DeleteEntryByKey(ctx context.Context, key string) error {
	return s(ctx, key)
}

// CreateStoreWith returns a functional store that can be used to conditionally save a types.KeyValueEntry.
// If the entry already exists, the item will not be updated an no error returned.
func (s *KeyValueStore) CreateStoreWith(client PutItemAPI) createKeyStore {
	return createKeyStore(func(ctx context.Context, entry *types.KeyValueEntry) error {
		record := &keyValueDynamoRecord{
			PrimaryKey: fmt.Sprintf("KEY#%s", entry.Key),
			Key:        entry.Key,
			Value:      entry.Value,
			CreatedAt:  time.Now(),
		}
		s.logger.Info(ctx, "saving record", slog.F("record", *record))

		input, err := createInputFromRecord(s.tableName, record)
		if err != nil {
			return err
		}

		s.logger.Debug(ctx, "input", slog.F("Item", input.Item))
		s.logger.Info(ctx, "conditional put to dynamodb")
		_, err = client.PutItem(ctx, input)
		var cfe *dbtypes.ConditionalCheckFailedException
		if errors.As(err, &cfe) {
			s.logger.Info(ctx, "condition failed", slog.Error(cfe))
			return nil
		}
		return err
	})
}

// GetStoreWith returns a functional store that can be used to retrieve a types.KeyValueEntry for a given key.
// If the entry is not found, no error is returned, and types.KeyValueEntry is nil.
func (s *KeyValueStore) GetStoreWith(client GetItemAPI) getKeyStore {
	return getKeyStore(func(ctx context.Context, key string) (*types.KeyValueEntry, error) {
		pk := fmt.Sprintf("KEY#%s", key)
		ctx = slog.With(ctx, slog.F("primary_key", pk))

		s.logger.Info(ctx, "retrieving record")

		keyAttr, err := attributevalue.MarshalMap(map[string]string{"pk": pk})
		if err != nil {
			return nil, err
		}

		output, err := client.GetItem(ctx, &dynamodb.GetItemInput{
			TableName: aws.String(s.tableName),
			Key:       keyAttr,
		})
		if err != nil {
			return nil, err
		}
		if len(output.Item) == 0 {
			s.logger.Warn(ctx, "record not found")
			return nil, nil
		}

		var record *keyValueDynamoRecord
		if err := attributevalue.UnmarshalMap(output.Item, &record); err != nil {
			return nil, err
		}
		s.logger.Debug(ctx, "found record", slog.F("record", *record))

		return &types.KeyValueEntry{
			Key:   record.Key,
			Value: record.Value,
		}, nil
	})
}

func (s *KeyValueStore) DeleteStoreWith(client DeleteItemAPI) deleteKeyStore {
	return deleteKeyStore(func(ctx context.Context, key string) error {
		pk := fmt.Sprintf("KEY#%s", key)
		ctx = slog.With(ctx, slog.F("primary_key", pk))

		s.logger.Info(ctx, "deleting record")

		keyAttr, err := attributevalue.MarshalMap(map[string]string{"pk": pk})
		if err != nil {
			return err
		}

		_, err = client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
			TableName: aws.String(s.tableName),
			Key:       keyAttr,
		})

		return err
	})
}

func createInputFromRecord(tableName string, record *keyValueDynamoRecord) (*dynamodb.PutItemInput, error) {
	cond := expression.AttributeNotExists(expression.Name("pk"))
	expr, err := expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		return nil, err
	}
	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return nil, err
	}

	return &dynamodb.PutItemInput{
		TableName:                 aws.String(tableName),
		Item:                      item,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	}, nil
}

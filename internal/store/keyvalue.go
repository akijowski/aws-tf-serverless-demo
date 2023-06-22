package store

import (
	"context"
	"fmt"
	"time"

	"cdr.dev/slog"
	"github.com/akijowski/aws-tf-serverless-demo/internal/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type PutItemAPI interface {
	PutItem(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type keyValueDynamoRecord struct {
	PrimaryKey string    `json:"pk"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at"`
}

type KeyValueStore struct {
	logger    slog.Logger
	tableName string
}

func With(logger slog.Logger, tableName string) *KeyValueStore {
	return &KeyValueStore{
		logger,
		tableName,
	}
}

func (s *KeyValueStore) CreateIfNotExists(ctx context.Context, client PutItemAPI, entry *types.KeyValueEntry) error {
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

	s.logger.Info(ctx, "conditional put to dynamodb")
	_, err = client.PutItem(ctx, input)

	return err
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

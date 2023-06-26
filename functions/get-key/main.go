package main

import (
	"context"
	"fmt"
	"os"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/sloghuman"
	"cdr.dev/slog/sloggers/slogjson"
	"github.com/akijowski/aws-tf-serverless-demo/internal/dynamo"
	"github.com/akijowski/aws-tf-serverless-demo/internal/store"
	"github.com/akijowski/aws-tf-serverless-demo/internal/tasks"
	"github.com/aws/aws-lambda-go/lambda"
)

const envTableName = "KV_TABLE_NAME"

func main() {
	ctx := context.Background()

	s := slogjson.Sink(os.Stdout)
	if isHumanLogs := os.Getenv("HUMAN_LOGS"); isHumanLogs != "" {
		s = sloghuman.Sink(os.Stdout)
	}
	logger := slog.Make(s).Named("get-key")

	dynamoClient, err := dynamo.NewClient(ctx, logger)
	if err != nil {
		logger.Fatal(context.Background(), "error building dynamo client", slog.Error(err))
	}
	tableName := os.Getenv(envTableName)
	if tableName == "" {
		logger.Fatal(ctx, fmt.Sprintf("missing required env: %q", envTableName))
	}

	kvStore := store.With(logger, tableName)
	task := tasks.NewGetKeyEntry(logger, kvStore.GetStoreWith(dynamoClient))

	lambda.Start(task.HandleGetKeyAPIRequest)
}

package main

import (
	"context"
	"os"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogjson"
	"github.com/akijowski/aws-tf-serverless-demo/internal/dynamo"
	"github.com/akijowski/aws-tf-serverless-demo/internal/store"
	"github.com/akijowski/aws-tf-serverless-demo/internal/tasks"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	s := slogjson.Sink(os.Stdout)
	logger := slog.Make(s).Named("get-key")

	dynamoClient, err := dynamo.NewClient()
	if err != nil {
		logger.Fatal(context.Background(), "error building dynamo client", slog.Error(err))
	}

	kvStore := store.With(logger, "my-table")
	task := tasks.NewGetKeyEntry(logger, kvStore.GetStoreWith(dynamoClient))

	lambda.Start(task.HandleGetKeyAPIRequest)
}

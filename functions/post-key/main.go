package main

import (
	"context"
	"os"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogjson"
	"github.com/akijowski/aws-tf-serverless-demo/internal/tasks"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	s := slogjson.Sink(os.Stdout)
	logger := slog.Make(s).Named("post-key")

	dynamoClient, err := newDynamoClient()
	if err != nil {
		logger.Fatal(context.Background(), "error building dynamo client", slog.F("error", err))
	}

	task := tasks.NewCreateKeyEntry(logger, "my-table", dynamoClient)

	lambda.Start(task.HandleCreateKeyAPIEvent)
}

func newDynamoClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

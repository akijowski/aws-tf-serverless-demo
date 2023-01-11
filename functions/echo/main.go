package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

var logger *log.Logger

func handler(ctx context.Context, event map[string]any) error {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		logger.Printf("%+v\n", lc)
	}
	logger.Printf("%+v\n", event)
	return nil
}

func main() {
	logger = log.Default()
	logger.SetPrefix("echo_lambda ")
	logger.SetFlags(log.Lshortfile | log.Lmsgprefix)

	lambda.Start(handler)
}

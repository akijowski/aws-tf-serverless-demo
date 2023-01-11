package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

var logger *log.Logger

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		logger.Printf("%+v\n", lc)
	}
	logger.Printf("%+v\n", req)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       `{"message": "hello world"}`,
	}, nil
}

func main() {
	logger = log.Default()
	logger.SetPrefix("hello_lambda ")
	logger.SetFlags(log.Lshortfile | log.Lmsgprefix)

	lambda.Start(handler)
}

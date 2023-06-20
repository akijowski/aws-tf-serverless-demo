package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

var logger *log.Logger

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func handler(ctx context.Context, req Request) (Response, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		logger.Printf("%+v\n", lc)
	}
	logger.Printf("%+v\n", req)
	name := "world"
	if nameParam, ok := req.QueryStringParameters["name"]; ok {
		name = nameParam
	}
	return Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf(`{"message": "hello %s"}`, name),
	}, nil
}

func main() {
	logger = log.Default()
	logger.SetPrefix("hello_lambda ")
	logger.SetFlags(log.Lshortfile | log.Lmsgprefix)

	lambda.Start(handler)
}

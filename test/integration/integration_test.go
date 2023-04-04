//go:build integration
// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCanInvokeHelloLambda(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	client := lambda.NewFromConfig(cfg, lambda.WithEndpointResolver(lambda.EndpointResolverFromURL("http://hello-lambda:8080")))

	out, err := client.Invoke(context.Background(), &lambda.InvokeInput{
		// The Lambda RIE does not support custom function names (yet)
		// https://github.com/aws/aws-lambda-runtime-interface-emulator/
		FunctionName: aws.String("function"),
	})

	if err != nil {
		t.Error(err)
	}
	if out.FunctionError != nil {
		t.Errorf("%s\n", aws.ToString(out.FunctionError))
	}
	t.Logf("%s\n", out.Payload)
}

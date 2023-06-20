//go:build integration
// +build integration

package integration

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestHelloLambda(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	client := lambda.NewFromConfig(cfg, lambda.WithEndpointResolver(lambda.EndpointResolverFromURL("http://hello-lambda:8080")))

	cases := map[string]struct {
		given events.APIGatewayProxyRequest
		want  string
	}{
		"empty query param returns default": {
			given: events.APIGatewayProxyRequest{
				Path: "/hello",
			},
			want: "{\"message\": \"hello world\"}",
		},
		"query param returns value": {
			given: events.APIGatewayProxyRequest{
				Path:                  "/hello",
				QueryStringParameters: map[string]string{"name": "joe"},
			},
			want: "{\"message\": \"hello joe\"}",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			tt := tt

			b, err := json.Marshal(tt.given)
			if err != nil {
				t.Fatal(err)
			}

			out, err := client.Invoke(context.Background(), &lambda.InvokeInput{
				// The Lambda RIE does not support custom function names (yet)
				// https://github.com/aws/aws-lambda-runtime-interface-emulator/
				FunctionName: aws.String("function"),
				Payload:      b,
			})
			if err != nil {
				t.Error(err)
			}
			if out.FunctionError != nil {
				t.Errorf("%s\n", aws.ToString(out.FunctionError))
			}

			var resp events.APIGatewayProxyResponse
			if err = json.Unmarshal(out.Payload, &resp); err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != 200 {
				t.Errorf("status code: %d\n", resp.StatusCode)
			}
			if resp.Body != tt.want {
				t.Error("incorrect body in response")
				t.Logf("wanted: %s\n", tt.want)
				t.Logf("got: %s\n", resp.Body)
			}

		})
	}
}

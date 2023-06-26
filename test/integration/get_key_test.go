//go:build integration
// +build integration

package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetKeyLambda(t *testing.T) {

	ctx := context.Background()

	testTableName := "get-key-values"

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithLogConfigurationWarnings(true),
		config.WithClientLogMode(aws.LogResponseWithBody))
	require.NoError(t, err)

	// create local dynamo
	localDynamo := newLocalDynamo(testTableName, KeyValueSchema)

	// create lambda client
	endpointOpt := lambda.WithEndpointResolver(lambda.EndpointResolverFromURL("http://get-key-lambda:8080"))
	lambdaClient := lambda.NewFromConfig(cfg, endpointOpt)

	cases := map[string]struct {
		given  events.APIGatewayProxyRequest
		dbData []map[string]types.AttributeValue
		want   events.APIGatewayProxyResponse
	}{
		"when entry exists returns ok": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				HTTPMethod:     http.MethodGet,
				Path:           "/keys/{keyID}",
				PathParameters: map[string]string{"keyID": "12345"},
			},
			dbData: []map[string]types.AttributeValue{
				{
					"pk":    &types.AttributeValueMemberS{Value: "KEY#12345"},
					"key":   &types.AttributeValueMemberS{Value: "12345"},
					"value": &types.AttributeValueMemberS{Value: "hijklmnop"},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"key":"12345","value":"hijklmnop"}`,
			},
		},
		"missing entry returns not found": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				HTTPMethod:     http.MethodGet,
				Path:           "/keys/{keyID}",
				PathParameters: map[string]string{"keyID": "12345"},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"message":"entry not found"}`,
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			tt := tt

			dynamoClient := localDynamo.newTest(t, cfg)

			if tt.dbData != nil && len(tt.dbData) > 0 {
				err := batchWrite(ctx, dynamoClient, testTableName, tt.dbData)
				require.NoError(t, err)
			}

			b, err := json.Marshal(tt.given)
			require.NoError(t, err)

			resp, err := lambdaClient.Invoke(ctx, &lambda.InvokeInput{
				FunctionName: aws.String("function"),
				Payload:      b,
			})

			var actual events.APIGatewayProxyResponse
			err = json.Unmarshal(resp.Payload, &actual)
			require.NoError(t, err)

			assert.Equal(t, tt.want, actual)
		})
	}
}

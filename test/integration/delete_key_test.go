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

func TestDeleteKeyLambda(t *testing.T) {
	ctx := context.Background()

	testTableName := "delete-key-values"

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithLogConfigurationWarnings(true),
		config.WithClientLogMode(aws.LogResponseWithBody))
	require.NoError(t, err)

	// create local dynamo
	localDynamo := newLocalDynamo(testTableName, KeyValueSchema)

	// create lambda client
	endpointOpt := lambda.WithEndpointResolver(lambda.EndpointResolverFromURL("http://delete-key-lambda:8080"))
	lambdaClient := lambda.NewFromConfig(cfg, endpointOpt)

	cases := map[string]struct {
		given  events.APIGatewayProxyRequest
		dbData []map[string]types.AttributeValue
		want   events.APIGatewayProxyResponse
	}{
		"when entry exists returns no content": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				HTTPMethod:     http.MethodDelete,
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
				StatusCode: http.StatusNoContent,
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
		},
		"missing entry returns no content": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				HTTPMethod:     http.MethodDelete,
				Path:           "/keys/{keyID}",
				PathParameters: map[string]string{"keyID": "12345"},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusNoContent,
				Headers:    map[string]string{"Content-Type": "application/json"},
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

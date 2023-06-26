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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTableName = "key-values"

func KeyValueSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(testTableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}
}

type batchWriteItemAPI interface {
	BatchWriteItem(context.Context, *dynamodb.BatchWriteItemInput, ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error)
}

func TestPostKeyLambda(t *testing.T) {

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithLogConfigurationWarnings(true),
		config.WithClientLogMode(aws.LogResponseWithBody))
	require.NoError(t, err)

	// create local dynamo
	localDynamo := newLocalDynamo(KeyValueSchema)

	// create lambda client
	endpointOpt := lambda.WithEndpointResolver(lambda.EndpointResolverFromURL("http://post-key-lambda:8080"))
	lambdaClient := lambda.NewFromConfig(cfg, endpointOpt)

	cases := map[string]struct {
		given  events.APIGatewayProxyRequest
		dbData []map[string]types.AttributeValue
		want   events.APIGatewayProxyResponse
	}{
		"when entry does not exist returns created": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				Body: `{"key":"abcdefg","value":"hijklmnop"}`,
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusCreated,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"key":"abcdefg","value":"hijklmnop"}`,
			},
		},
		"when entry exists returns created": {
			given: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					RequestID: "aws-abc-1234567890",
				},
				Body: `{"key":"12345","value":"hijklmnop"}`,
			},
			dbData: []map[string]types.AttributeValue{
				{
					"pk":    &types.AttributeValueMemberS{Value: "KEY#12345"},
					"key":   &types.AttributeValueMemberS{Value: "12345"},
					"value": &types.AttributeValueMemberS{Value: "hijklmnop"},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusCreated,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"key":"12345","value":"hijklmnop"}`,
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

type localDynamo struct {
	schema func() *dynamodb.CreateTableInput
}

func newLocalDynamo(schema func() *dynamodb.CreateTableInput) *localDynamo {
	return &localDynamo{
		schema: schema,
	}
}

func (l *localDynamo) newTest(t testing.TB, cfg aws.Config) *dynamodb.Client {
	ctx := context.Background()
	endpointOpt := dynamodb.WithEndpointResolver(dynamodb.EndpointResolverFromURL("http://dynamodb-local:8000"))
	cfg = cfg.Copy()
	cfg.ClientLogMode = 0
	client := dynamodb.NewFromConfig(cfg, endpointOpt)

	schema := l.schema()

	_, err := client.CreateTable(ctx, schema)
	if err != nil {
		t.Fatalf("unable to create table: %v\n", err)
	}

	t.Cleanup(func() {
		_, err = client.DeleteTable(ctx, &dynamodb.DeleteTableInput{
			TableName: schema.TableName,
		})
		if err != nil {
			t.Fatalf("unable to delete table: %v\n", err)
		}
	})

	return client
}

func batchWrite(ctx context.Context, client batchWriteItemAPI, tableName string, items []map[string]types.AttributeValue) error {
	writes := make([]types.WriteRequest, 0)
	for _, item := range items {
		write := types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		}
		writes = append(writes, write)
	}

	_, err := client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			tableName: writes,
		},
	})

	return err
}

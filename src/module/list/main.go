package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/internal/create"
	"github.com/brittandeyoung/tfregistry/src/internal/resource/module/odm"
)

var ddb dynamodb.Client
var table string

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	ddb = *dynamodb.NewFromConfig(sdkConfig)
	table = os.Getenv("table_name")
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	namespace, ok := req.PathParameters["namespace"]
	if !ok {
		return create.ClientError(http.StatusBadRequest)
	}

	module := odm.Module{
		Namespace:    namespace,
		ResourceType: odm.DynamoDbType,
	}

	provider, ok := req.QueryStringParameters["provider"]
	if ok {
		module.Provider = provider
	}

	items, err := module.List(ctx, ddb, table)

	if err != nil {
		return create.ServerError(err)
	}

	if items == nil {
		return create.ClientError(http.StatusNotFound)
	}

	json, err := json.Marshal(items)
	if err != nil {
		return create.ServerError(err)
	}
	log.Printf("Successfully fetched item %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(handler)
}

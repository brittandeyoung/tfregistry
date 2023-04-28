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

	provider, ok := req.PathParameters["provider"]
	if !ok {
		return create.ClientError(http.StatusBadRequest)
	}

	name, ok := req.PathParameters["name"]
	if !ok {
		return create.ClientError(http.StatusBadRequest)
	}
	var module odm.Module

	json.Unmarshal([]byte(req.Body), &module)

	module.Namespace = namespace
	module.Provider = provider
	module.Name = name

	item, err := module.Update(ctx, ddb, table)

	if err != nil {
		return create.ServerError(err)
	}

	json, err := json.Marshal(item)

	if err != nil {
		return create.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(handler)
}

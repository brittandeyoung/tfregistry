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
	"github.com/brittandeyoung/tfregistry/src/api/internal/create"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/namespace/odm"
	"github.com/brittandeyoung/tfregistry/src/api/internal/validate"
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
	var namespace odm.Namespace
	json.Unmarshal([]byte(req.Body), &namespace)

	item, err := namespace.Create(ctx, ddb, table)

	if validate.ConditionalCheckFailedException(err) {
		return create.ServerErrorConflict(err)
	}

	if err != nil {
		return create.ServerError(err)
	}

	json, err := json.Marshal(item)

	if err != nil {
		return create.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(handler)
}

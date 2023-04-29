package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/create"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/namespace/odm"
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
	name, ok := req.PathParameters["namespace"]
	if !ok {
		return create.ClientError(http.StatusBadRequest)
	}

	namespace := odm.Namespace{
		Name:         name,
		ResourceType: odm.DynamoDbType,
	}

	err := namespace.Delete(ctx, ddb, table)

	if err != nil {
		return create.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Body:       "",
	}, nil
}

func main() {
	lambda.Start(handler)
}

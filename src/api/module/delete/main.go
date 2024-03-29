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
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/module"
)

type deps struct {
	ddb   ddb.DynamoDeleteItemAPI
	table string
}

func (d *deps) handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if d.ddb == nil {
		sdkConfig, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatal(err)
		}

		table := os.Getenv("TABLE_NAME")

		d.ddb = dynamodb.NewFromConfig(sdkConfig)
		d.table = table
	}

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

	in := &module.DeleteModuleInput{
		Pk: module.FlattenPartitionKey(namespace),
		Sk: module.FlattenSortKey(namespace, provider, name),
	}

	err := module.Delete(ctx, d.ddb, d.table, in)

	if err != nil {
		return create.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Body:       "",
	}, nil
}

func main() {
	d := deps{}
	lambda.Start(d.handler)
}

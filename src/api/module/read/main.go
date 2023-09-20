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
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/module"
)

type deps struct {
	ddb   ddb.DynamoGetItemAPI
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

	in := &module.GetModuleInput{
		Pk: module.FlattenPartitionKey(namespace),
		Sk: module.FlattenSortKey(namespace, provider, name),
	}

	item, err := module.Read(ctx, d.ddb, d.table, in)

	if err != nil {
		return create.ServerError(err)
	}

	if item == nil {
		return create.ClientError(http.StatusNotFound)
	}

	json, err := json.Marshal(item)
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
	d := deps{}
	lambda.Start(d.handler)
}

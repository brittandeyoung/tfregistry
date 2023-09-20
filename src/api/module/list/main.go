package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/create"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/module"
)

type deps struct {
	ddb   ddb.DynamoQueryAPI
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

	var startKey *string
	startKeyString, ok := req.QueryStringParameters["start_key"]
	if ok {
		startKey = aws.String(startKeyString)
	}

	var limit *int32
	limitString, ok := req.QueryStringParameters["limit"]
	if ok {
		limitInt, err := strconv.Atoi(limitString)
		if err != nil {
			return create.ClientError(http.StatusBadRequest)
		}
		limit = aws.Int32(int32(limitInt))
	}

	in := &module.ListModuleInput{
		Namespace: namespace,
		StartKey:  startKey,
		Limit:     limit,
	}

	provider, ok := req.QueryStringParameters["provider"]
	if ok {
		in.Provider = provider
	}

	name, ok := req.QueryStringParameters["name"]
	if ok {
		in.Name = name
	}

	out, err := module.List(ctx, d.ddb, d.table, in)

	if err != nil {
		return create.ServerError(err)
	}

	json, err := json.Marshal(out)

	if err != nil {
		return create.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func main() {
	d := deps{}
	lambda.Start(d.handler)
}

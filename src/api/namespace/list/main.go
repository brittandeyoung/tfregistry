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
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/namespace"
)

type deps struct {
	ddb                        ddb.DynamoQueryAPI
	table                      string
	AccessControlAllowedHeader string
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

	if d.AccessControlAllowedHeader == "" {
		d.AccessControlAllowedHeader = os.Getenv("ACCESS_CONTROL_ALLOWED_HEADER")
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

	in := namespace.ListNamespaceInput{
		StartKey: startKey,
		Limit:    limit,
	}

	out, err := namespace.List(ctx, d.ddb, d.table, in)

	if err != nil {
		return create.ServerError(err)
	}

	json, err := json.Marshal(out)
	if err != nil {
		return create.ServerError(err)
	}
	log.Printf("Successfully fetched item %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": d.AccessControlAllowedHeader,
		},
	}, nil
}

func main() {
	d := deps{}
	lambda.Start(d.handler)
}

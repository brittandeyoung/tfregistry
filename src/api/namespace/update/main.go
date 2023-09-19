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
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/namespace"
)

type deps struct {
	ddb   ddb.DynamoUpdateItemAPI
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

	name, ok := req.PathParameters["namespace"]
	if !ok {
		return create.ClientError(http.StatusBadRequest)
	}

	in := &namespace.UpdateNamespaceInput{}
	in.Pk = namespace.Pk
	in.Sk = name
	json.Unmarshal([]byte(req.Body), in)

	item, err := namespace.Update(ctx, d.ddb, d.table, in)

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
	d := deps{}
	lambda.Start(d.handler)
}

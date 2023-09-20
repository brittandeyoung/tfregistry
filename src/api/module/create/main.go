package main

import (
	"context"
	"encoding/json"
	"errors"
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
	"github.com/brittandeyoung/tfregistry/src/api/internal/validate"
)

type deps struct {
	ddb   ddb.DynamoPutItemAPI
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

	in := &module.CreateModuleInput{}

	json.Unmarshal([]byte(req.Body), in)

	in.Namespace = namespace

	item, err := module.Create(ctx, d.ddb, d.table, in)

	if validate.ConditionalCheckFailedException(err) {
		return create.ServerErrorConflict(errors.New("a module with this name and provider combination already exists."))
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
	d := deps{}
	lambda.Start(d.handler)
}

package create

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func ClientError(status int) (events.APIGatewayProxyResponse, error) {
	errorMessage := ErrorMessage{
		Error: http.StatusText(status),
	}
	body, err := json.Marshal(errorMessage)
	if err != nil {
		return ServerError(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: status,
	}, nil
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func ClientErrorCustomMessage(status int, err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())
	errorMessage := ErrorMessage{
		Error: err.Error(),
	}
	body, err := json.Marshal(errorMessage)
	if err != nil {
		return ServerError(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: status,
	}, nil
}

func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       "{\"error\": \"" + http.StatusText(http.StatusInternalServerError) + "\"}",
		StatusCode: http.StatusInternalServerError,
	}, nil
}

func ServerErrorConflict(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())
	errorMessage := ErrorMessage{
		Error: err.Error(),
	}
	body, err := json.Marshal(errorMessage)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusConflict,
	}, nil
}

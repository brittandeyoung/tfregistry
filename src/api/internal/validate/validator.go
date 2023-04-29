package validate

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ConditionalCheckFailedException(err error) bool {
	var e *types.ConditionalCheckFailedException
	return errors.As(err, &e)
}

package odm

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ConditionalCheckFailedException(err error) bool {
	var e *types.ConditionalCheckFailedException
	return errors.As(err, &e)
}

func ValidateRequiredFields(m *Module) error {
	if m.Namespace == "" || m.Provider == "" || m.Name == "" {
		return errors.New("module is missing one of the required fields (Namespace, Provider, or Name)")
	}

	return nil
}

func ValidateListFields(m *Module) error {
	if m.Namespace == "" {
		return errors.New("list module is missing the required field (Namespace)")
	}

	return nil
}

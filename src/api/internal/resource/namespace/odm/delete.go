package odm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (m *Namespace) Delete(ctx context.Context, ddb dynamodb.Client, table string) error {
	err := ValidateRequiredFields(m)

	if err != nil {
		return err
	}

	key, err := m.ExpandPartitionKeyAndSortKey()

	if err != nil {
		return err
	}

	in := &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       key,
	}

	_, err = ddb.DeleteItem(ctx, in)

	if err != nil {
		return err
	}

	return nil
}

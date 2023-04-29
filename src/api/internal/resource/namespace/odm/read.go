package odm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (m *Namespace) Read(ctx context.Context, ddb dynamodb.Client, table string) (*Namespace, error) {
	err := ValidateRequiredFields(m)

	if err != nil {
		return nil, err
	}

	key, err := m.ExpandPartitionKeyAndSortKey()

	if err != nil {
		return nil, err
	}

	in := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       key,
	}

	result, err := ddb.GetItem(ctx, in)

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	item := new(Namespace)
	err = attributevalue.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

package namespace

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type GetNamespaceInput struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

func Read(ctx context.Context, ddbClient ddb.DynamoGetItemAPI, table string, m GetNamespaceInput) (*Namespace, error) {
	key, err := attributevalue.MarshalMap(m)

	if err != nil {
		return nil, err
	}

	in := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       key,
	}

	result, err := ddbClient.GetItem(ctx, in)

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

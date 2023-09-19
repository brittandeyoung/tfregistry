package namespace

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type DeleteNamespaceInput struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

func Delete(ctx context.Context, ddb ddb.DynamoDeleteItemAPI, table string, m DeleteNamespaceInput) error {

	key, err := attributevalue.MarshalMap(m)

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

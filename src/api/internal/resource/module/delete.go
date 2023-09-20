package module

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type DeleteModuleInput struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

func Delete(ctx context.Context, ddbClient ddb.DynamoDeleteItemAPI, table string, m *DeleteModuleInput) error {
	key, err := ExpandPartitionKeyAndSortKey(m.Pk, m.Sk)

	if err != nil {
		return err
	}

	in := &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       key,
	}

	_, err = ddbClient.DeleteItem(ctx, in)

	if err != nil {
		return err
	}

	return nil
}

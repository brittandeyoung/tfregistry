package module

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type UpdateModuleInput struct {
	Pk          string  `json:"pk" dynamodbav:"pk"`
	Sk          string  `json:"sk" dynamodbav:"sk"`
	Description *string `json:"description" dynamodbav:"description"`
	Source      *string `json:"source" dynamodbav:"source"`
	Verified    *bool   `json:"verified" dynamodbav:"verified"`
}

func Update(ctx context.Context, ddbClient ddb.DynamoUpdateItemAPI, table string, m *UpdateModuleInput) (*Module, error) {
	key, err := ExpandPartitionKeyAndSortKey(m.Pk, m.Sk)

	if err != nil {
		return nil, err
	}

	update := expression.Set(expression.Name("description"), expression.Value(m.Description)).
		Set(expression.Name("source"), expression.Value(m.Source)).Set(expression.Name("verified"), expression.Value(m.Verified))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		return nil, err
	}

	in := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table),
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	}

	res, err := ddbClient.UpdateItem(ctx, in)

	if err != nil {
		return nil, err
	}

	item := new(Module)
	err = attributevalue.UnmarshalMap(res.Attributes, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

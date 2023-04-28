package odm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (m *Module) Create(ctx context.Context, ddb dynamodb.Client, table string) (*Module, error) {
	err := ValidateRequiredFields(m)

	if err != nil {
		return nil, err
	}

	m.Id = FlattenId(m.Namespace, m.Name, m.Provider)
	m.SortKey = FlattenSortKey(m.Namespace, m.Provider, m.Name)
	m.ResourceType = DynamoDbType

	item, err := attributevalue.MarshalMap(m)

	if err != nil {
		return nil, err
	}

	primaryKeyCheckCondition := expression.Name("resourceType").NotEqual(expression.Value(m.ResourceType)).And(expression.Name("sortKey").NotEqual(expression.Value(m.SortKey)))
	conditionBuilder := expression.NewBuilder().WithCondition(primaryKeyCheckCondition)
	expr, _ := conditionBuilder.Build()

	in := &dynamodb.PutItemInput{
		TableName:                 aws.String(table),
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
	}

	res, err := ddb.PutItem(ctx, in)

	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(res.Attributes, &item)
	if err != nil {
		return nil, err
	}

	return m, nil
}

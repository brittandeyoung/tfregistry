package odm

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (m *Module) List(ctx context.Context, ddb dynamodb.Client, table string) ([]*Module, error) {
	err := ValidateListFields(m)

	if err != nil {
		return nil, err
	}

	condition := expression.Name("resourceType").Equal(expression.Value(m.ResourceType)).And(expression.Name("sortKey").BeginsWith(m.FlattenQuerySortKey()))
	condExp, err := expression.NewBuilder().WithCondition(condition).Build()

	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	}

	in := &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		KeyConditionExpression:    condExp.Condition(),
		ExpressionAttributeNames:  condExp.Names(),
		ExpressionAttributeValues: condExp.Values(),
	}

	result, err := ddb.Query(ctx, in)

	if err != nil {
		return nil, err
	}

	items := make([]*Module, 0)

	if len(result.Items) == 0 {
		return nil, nil
	}

	for _, resultItem := range result.Items {
		item := new(Module)
		err = attributevalue.UnmarshalMap(resultItem, item)

		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}

	return items, nil
}

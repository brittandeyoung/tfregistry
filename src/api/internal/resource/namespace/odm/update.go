package odm

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (m *Namespace) Update(ctx context.Context, ddb dynamodb.Client, table string) (*Namespace, error) {
	err := ValidateRequiredFields(m)

	if err != nil {
		return nil, err
	}

	if m.Description == "" && m.Email == "" {
		return nil, errors.New("updating a module requires an updated description or source")
	}

	key, err := m.ExpandPartitionKeyAndSortKey()

	if err != nil {
		return nil, err
	}

	update := expression.Set(expression.Name("description"), expression.Value(m.Description)).
		Set(expression.Name("email"), expression.Value(m.Email))

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
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	res, err := ddb.UpdateItem(ctx, in)

	if err != nil {
		return nil, err
	}

	item := new(Namespace)
	err = attributevalue.UnmarshalMap(res.Attributes, item)

	if err != nil {
		return nil, err
	}

	return m, nil
}

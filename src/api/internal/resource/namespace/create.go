package namespace

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type CreateNamespaceInput struct {
	Id          string  `json:"id" dynamodbav:"id"`
	Pk          string  `json:"pk" dynamodbav:"pk"`
	Sk          string  `json:"sk" dynamodbav:"sk"`
	Description *string `json:"description" dynamodbav:"description"`
	Email       *string `json:"email" dynamodbav:"email"`
	Name        string  `json:"name" dynamodbav:"name"`
}

func Create(ctx context.Context, ddbClient ddb.DynamoPutItemAPI, table string, m *CreateNamespaceInput) (*Namespace, error) {
	if m.Name == "" {
		return nil, errors.New("namespace is missing one of the required fields (Name)")
	}

	m.Id = m.Name
	m.Sk = m.Name
	m.Pk = Pk

	item, err := attributevalue.MarshalMap(m)

	if err != nil {
		return nil, err
	}

	primaryKeyCheckCondition := expression.Name("pk").NotEqual(expression.Value(m.Pk)).And(expression.Name("sk").NotEqual(expression.Value(m.Sk)))
	conditionBuilder := expression.NewBuilder().WithCondition(primaryKeyCheckCondition)
	expr, _ := conditionBuilder.Build()

	in := &dynamodb.PutItemInput{
		TableName:                 aws.String(table),
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
	}

	_, err = ddbClient.PutItem(ctx, in)

	if err != nil {
		return nil, err
	}

	out := &Namespace{
		Id:          m.Id,
		Description: m.Description,
		Email:       m.Email,
		Name:        m.Name,
	}

	return out, nil
}

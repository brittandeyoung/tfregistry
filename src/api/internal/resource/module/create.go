package module

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type CreateModuleInput struct {
	Id          string         `json:"id" dynamodbav:"id"`
	Pk          string         `json:"pk" dynamodbav:"pk"`
	Sk          string         `json:"sk" dynamodbav:"sk"`
	Description *string        `json:"description" dynamodbav:"description"`
	Downloads   *int           `json:"downloads" dynamodbav:"downloads"`
	Name        string         `json:"name" dynamodbav:"name"`
	Namespace   string         `json:"namespace" dynamodbav:"namespace"`
	Provider    string         `json:"provider" dynamodbav:"provider"`
	Source      *string        `json:"source" dynamodbav:"source"`
	Verified    *bool          `json:"verified" dynamodbav:"verified"`
	Versions    *[]interface{} `json:"versions" dynamodbav:"versions"`
}

func Create(ctx context.Context, ddbClient ddb.DynamoPutItemAPI, table string, m *CreateModuleInput) (*Module, error) {
	if m.Namespace == "" || m.Provider == "" || m.Name == "" {
		return nil, errors.New("module is missing one of the required fields (Namespace, Provider, or Name)")
	}

	m.Id = FlattenId(m.Namespace, m.Name, m.Provider)
	m.Sk = FlattenSortKey(m.Namespace, m.Provider, m.Name)
	m.Pk = FlattenPartitionKey(m.Namespace)

	item, err := attributevalue.MarshalMap(m)

	if err != nil {
		return nil, err
	}

	in := &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(pk) AND attribute_not_exists(sk)"),
	}

	res, err := ddbClient.PutItem(ctx, in)

	if err != nil {
		return nil, err
	}

	out := &Module{
		Id:          m.Id,
		Description: m.Description,
		Downloads:   m.Downloads,
		Name:        m.Name,
		Namespace:   m.Namespace,
		Provider:    m.Provider,
		Source:      m.Source,
		Verified:    m.Verified,
	}

	err = attributevalue.UnmarshalMap(res.Attributes, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

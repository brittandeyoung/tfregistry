package version

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type CreateModuleVersionInput struct {
	Pk           string                    `json:"pk" dynamodbav:"pk"`
	Sk           string                    `json:"sk" dynamodbav:"sk"`
	Id           string                    `json:"id" dynamodbav:"id"`
	Code         *string                   `json:"code" dynamodbav:"code"`
	Dependencies Dependencies              `json:"dependencies" dynamodbav:"dependencies"`
	Downloads    *int                      `json:"downloads" dynamodbav:"downloads"`
	Module       string                    `json:"module" dynamodbav:"module"`
	Outputs      *[]map[string]interface{} `json:"outputs" dynamodbav:"outputs"`
	Readme       *string                   `json:"readme" dynamodbav:"readme"`
	Resources    *[]map[string]interface{} `json:"resource" dynamodbav:"resource"`
	Version      string                    `json:"version" dynamodbav:"version"`
	Variables    *[]map[string]interface{} `json:"variables" dynamodbav:"variables"`
}

func Create(ctx context.Context, ddbClient ddb.DynamoPutItemAPI, table string, m *CreateModuleVersionInput) (*ModuleVersion, error) {
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

	out := &ModuleVersion{
		Id:           m.Id,
		Dependencies: m.Dependencies,
		Module:       m.Module,
		Outputs:      m.Outputs,
		Readme:       m.Readme,
		Resources:    m.Resources,
		Version:      m.Version,
		Variables:    m.Variables,
	}

	return out, nil
}

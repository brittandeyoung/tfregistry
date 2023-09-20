package namespace

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type UpdateNamespaceInput struct {
	Pk          string  `json:"pk" dynamodbav:"pk"`
	Sk          string  `json:"sk" dynamodbav:"sk"`
	Description *string `json:"description" dynamodbav:"description"`
	Email       *string `json:"email" dynamodbav:"email"`
}

func Update(ctx context.Context, ddbClient ddb.DynamoUpdateItemAPI, table string, m *UpdateNamespaceInput) (*Namespace, error) {
	if m.Description == nil && m.Email == nil {
		return nil, errors.New("updating a module requires an updated description or source")
	}

	itemKey := ddb.AttributeValueKey{
		Pk: m.Pk,
		Sk: m.Sk,
	}

	key, err := attributevalue.MarshalMap(itemKey)

	if err != nil {
		return nil, err
	}

	// Update attributes that "should" be udpated. Currently only the email and description attributes
	// The pk, sk, cidr, cidrType and location fields aren't updated because they compose the item's identity
	update := expression.Set(expression.Name("email"), expression.Value(m.Email))
	update.Set(expression.Name("description"), expression.Value(m.Description))
	condition := expression.AttributeExists(expression.Name("pk")).And(expression.AttributeExists(expression.Name("sk")))
	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(condition).Build()

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
		ConditionExpression:       expr.Condition(),
	}

	res, err := ddbClient.UpdateItem(ctx, in)

	if err != nil {
		return nil, err
	}

	item := new(Namespace)
	err = attributevalue.UnmarshalMap(res.Attributes, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

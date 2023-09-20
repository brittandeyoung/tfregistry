package module

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common"
	"github.com/brittandeyoung/tfregistry/src/api/internal/resource/common/ddb"
)

type ListModuleInput struct {
	Namespace string
	Provider  string
	Name      string
	Limit     *int32
	StartKey  *string
}

type ListModuleOutput struct {
	Meta    common.Meta
	Modules []*Module
}

func List(ctx context.Context, ddbClient ddb.DynamoQueryAPI, table string, m *ListModuleInput) (*ListModuleOutput, error) {
	condition := expression.Name("pk").Equal(expression.Value(FlattenPartitionKey(m.Namespace))).And(expression.Name("sk").BeginsWith(FlattenQuerySortKey(m.Namespace, m.Provider, m.Name)))
	condExp, err := expression.NewBuilder().WithCondition(condition).Build()

	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	}

	out := &ListModuleOutput{}
	in := &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		KeyConditionExpression:    condExp.Condition(),
		ExpressionAttributeNames:  condExp.Names(),
		ExpressionAttributeValues: condExp.Values(),
	}

	if m.Limit != nil {
		in.Limit = m.Limit
	}

	if m.StartKey != nil {
		startKeyObject := ddb.AttributeValueKey{
			Pk: Pk,
			Sk: aws.ToString(m.StartKey),
		}
		avStartKey, err := attributevalue.MarshalMap(startKeyObject)
		if err != nil {
			return nil, err
		}

		in.ExclusiveStartKey = avStartKey
	}

	result, err := ddbClient.Query(ctx, in)

	if err != nil {
		return nil, err
	}

	items := make([]*Module, 0)

	for _, resultItem := range result.Items {
		item := new(Module)
		err = attributevalue.UnmarshalMap(resultItem, item)

		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}

	fullUrl, err := url.Parse(fmt.Sprintf("/api/modules/%s", m.Namespace))

	if err != nil {
		return nil, err
	}

	values := fullUrl.Query()

	if m.Limit != nil {
		values.Add("limit", fmt.Sprintf("%d", int(aws.ToInt32(m.Limit))))
	}

	var lastEvalSortKey *string
	var nextUrl *string
	lek := map[string]string{}
	if result.LastEvaluatedKey != nil {
		attributevalue.UnmarshalMap(result.LastEvaluatedKey, &lek)
		lastEvalSortKey = aws.String(lek["sk"])
		values.Add("start_key", aws.ToString(lastEvalSortKey))
		fullUrl.RawQuery = values.Encode()
		nextUrl = aws.String(fullUrl.String())
	}

	out.Meta = common.Meta{
		Count:            result.Count,
		LastEvaluatedKey: lastEvalSortKey,
		NextUrl:          nextUrl,
	}
	out.Modules = items

	return out, nil
}

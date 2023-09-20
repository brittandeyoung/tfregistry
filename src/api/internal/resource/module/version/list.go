package version

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

type ListModuleVersionInput struct {
	Namespace string
	Provider  string
	Name      string
	Limit     *int32
	StartKey  *string
}

type ListModuleVersionOutput struct {
	Meta     common.Meta
	Versions []*ModuleVersion
}

func List(ctx context.Context, ddbClient ddb.DynamoQueryAPI, table string, m *ListModuleVersionInput) (*ListModuleVersionOutput, error) {
	condition := expression.Name("pk").Equal(expression.Value(FlattenPartitionKey(FlattenModule(m.Namespace, m.Name, m.Provider))))
	condExp, err := expression.NewBuilder().WithCondition(condition).Build()

	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	}

	out := &ListModuleVersionOutput{}
	in := &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		KeyConditionExpression:    condExp.Condition(),
		ExpressionAttributeNames:  condExp.Names(),
		ExpressionAttributeValues: condExp.Values(),
		ScanIndexForward:          aws.Bool(false),
	}

	result, err := ddbClient.Query(ctx, in)

	if err != nil {
		return nil, err
	}

	items := make([]*ModuleVersion, 0)

	if len(result.Items) == 0 {
		return nil, nil
	}

	for _, resultItem := range result.Items {
		item := new(ModuleVersion)
		err = attributevalue.UnmarshalMap(resultItem, item)

		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}

	fullUrl, err := url.Parse(fmt.Sprintf("/api/modules/%s/%s/%s", m.Namespace, m.Name, m.Provider))

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
	out.Versions = items

	return out, nil
}

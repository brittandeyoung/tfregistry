package module

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func FlattenId(namespace string, name string, provider string) string {
	return fmt.Sprintf("%s/%s/%s", namespace, name, provider)
}

func FlattenSortKey(namespace string, provider string, name string) string {
	return fmt.Sprintf("%s/%s/%s", namespace, provider, name)
}

func FlattenPartitionKey(namespace string) string {
	return namespace + "/" + Pk
}

func ExpandPartitionKeyAndSortKey(pk, sk string) (map[string]types.AttributeValue, error) {
	moduleKey := ModuleKey{
		Pk: pk,
		Sk: sk,
	}

	key, err := attributevalue.MarshalMap(moduleKey)

	if err != nil {
		return nil, err
	}

	return key, nil
}

func FlattenQuerySortKey(namespace, provider, name string) string {
	sortKey := namespace + "/"
	if provider != "" {
		sortKey = sortKey + provider + "/"
	}
	if name != "" && provider != "" {
		sortKey = sortKey + name
	}

	return sortKey
}

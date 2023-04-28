package odm

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

func (m *Module) ExpandPartitionKeyAndSortKey() (map[string]types.AttributeValue, error) {
	moduleKey := ModuleKey{
		SortKey:      FlattenSortKey(m.Namespace, m.Provider, m.Name),
		ResourceType: DynamoDbType,
	}

	key, err := attributevalue.MarshalMap(moduleKey)

	if err != nil {
		return nil, err
	}

	return key, nil
}

func (m *Module) FlattenQuerySortKey() string {
	sortKey := m.Namespace + "/"
	if m.Provider != "" {
		sortKey = sortKey + m.Provider + "/"
	}

	return sortKey
}

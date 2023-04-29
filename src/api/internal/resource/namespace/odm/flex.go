package odm

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (m *Namespace) ExpandPartitionKeyAndSortKey() (map[string]types.AttributeValue, error) {
	namespaceKey := NamespaceKey{
		SortKey:      m.Name,
		ResourceType: DynamoDbType,
	}

	key, err := attributevalue.MarshalMap(namespaceKey)

	if err != nil {
		return nil, err
	}

	return key, nil
}

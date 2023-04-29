package odm

const (
	DynamoDbType = "namespace"
)

type Namespace struct {
	Description  string `json:"description" dynamodbav:"description"`
	Id           string `json:"id" dynamodbav:"id"`
	Email        string `json:"email" dynamodbav:"email"`
	Name         string `json:"name" dynamodbav:"name"`
	SortKey      string `json:"sortKey" dynamodbav:"sortKey"`
	ResourceType string `json:"resourceType" dynamodbav:"resourceType"`
}

type NamespaceKey struct {
	SortKey      string `json:"sortKey" dynamodbav:"sortKey"`
	ResourceType string `json:"resourceType" dynamodbav:"resourceType"`
}

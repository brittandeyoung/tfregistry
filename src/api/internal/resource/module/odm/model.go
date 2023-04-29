package odm

const (
	DynamoDbType = "module"
)

type Module struct {
	Description  string `json:"description" dynamodbav:"description"`
	Id           string `json:"id" dynamodbav:"id"`
	Name         string `json:"name" dynamodbav:"name"`
	Namespace    string `json:"namespace" dynamodbav:"namespace"`
	Provider     string `json:"provider" dynamodbav:"provider"`
	SortKey      string `json:"sortKey" dynamodbav:"sortKey"`
	Source       string `json:"source" dynamodbav:"source"`
	ResourceType string `json:"resourceType" dynamodbav:"resourceType"`
}

type ModuleKey struct {
	SortKey      string `json:"sortKey" dynamodbav:"sortKey"`
	ResourceType string `json:"resourceType" dynamodbav:"resourceType"`
}

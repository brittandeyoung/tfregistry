package module

const (
	Pk = "module"
)

type Module struct {
	Id          string         `json:"id" dynamodbav:"id"`
	Description *string        `json:"description" dynamodbav:"description"`
	Downloads   *int           `json:"downloads" dynamodbav:"downloads"`
	Name        string         `json:"name" dynamodbav:"name"`
	Namespace   string         `json:"namespace" dynamodbav:"namespace"`
	Provider    string         `json:"provider" dynamodbav:"provider"`
	Source      *string        `json:"source" dynamodbav:"source"`
	Verified    *bool          `json:"verified" dynamodbav:"verified"`
	Versions    *[]interface{} `json:"versions" dynamodbav:"versions"`
}

type ModuleKey struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

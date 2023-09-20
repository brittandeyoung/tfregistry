package version

const (
	Pk = "module/version"
)

type ModuleVersion struct {
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

type ModuleVersionKey struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

type Dependencies struct {
	Providers []map[string]interface{} `json:"providers" dynamodbav:"providers"`
	Modules   []map[string]interface{} `json:"modules" dynamodbav:"modules"`
}

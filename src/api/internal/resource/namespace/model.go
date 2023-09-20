package namespace

const (
	Pk = "namespace"
)

type Namespace struct {
	Id          string  `json:"id" dynamodbav:"id"`
	Description *string `json:"description" dynamodbav:"description"`
	Email       *string `json:"email" dynamodbav:"email"`
	Name        string  `json:"name" dynamodbav:"name"`
}

package common

type SharedFields struct {
	Id string `json:"id" dynamodbav:"id"`
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

type Meta struct {
	Count            int32   `json:"count"`
	LastEvaluatedKey *string `json:"last_evaluated_key"`
	NextUrl          *string `json:"next_url"`
}

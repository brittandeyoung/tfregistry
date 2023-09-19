package ddb

type AttributeValueKey struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

package version

import (
	"fmt"
	"strings"
)

func FlattenId(module string, version string) string {
	return fmt.Sprintf("%s/%s", module, version)
}

// func ExpandModule(moduleId string) *moduleOdm.Module {
// 	idParts := strings.Split(moduleId, "/")

// 	module := new(moduleOdm.Module)
// 	module.Namespace = idParts[0]
// 	module.Name = idParts[1]
// 	module.Provider = idParts[2]

// 	return module
// }

func FlattenModule(namespace string, name string, provider string) string {
	return fmt.Sprintf("%s/%s/%s", namespace, name, provider)
}

func FlattenSortKey(module, version string) string {
	idParts := strings.Split(version, ".")

	major := fmt.Sprintf("%010s", idParts[0])
	minor := fmt.Sprintf("%010s", idParts[1])
	patch := fmt.Sprintf("%010s", idParts[2])
	return fmt.Sprintf("%s/%s.%s.%s", module, major, minor, patch)
}

func FlattenPartitionKey(module string) string {
	return Pk + "/" + module
}

// func ExpandPartitionKeyAndSortKey() (map[string]types.AttributeValue, error) {
// 	ModuleVersionKey := ModuleVersionKey{
// 		Sk: FlattenSortKey(m.Version),
// 		Pk: FlattenPartitionKey(m.Module),
// 	}

// 	key, err := attributevalue.MarshalMap(ModuleVersionKey)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }

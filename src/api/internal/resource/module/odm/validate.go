package odm

import (
	"errors"
)

func ValidateRequiredFields(m *Module) error {
	if m.Namespace == "" || m.Provider == "" || m.Name == "" {
		return errors.New("module is missing one of the required fields (Namespace, Provider, or Name)")
	}

	return nil
}

func ValidateListFields(m *Module) error {
	if m.Namespace == "" {
		return errors.New("list module is missing the required field (Namespace)")
	}

	return nil
}

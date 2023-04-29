package odm

import "errors"

func ValidateRequiredFields(m *Namespace) error {
	if m.Name == "" {
		return errors.New("module is missing one of the required fields (Namespace, Provider, or Name)")
	}

	return nil
}

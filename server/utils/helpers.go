package utils

import "fmt"

func CheckValidStringMap(m map[string]bool, prop string, customErrMsg string) (bool, error) {
	if _, exists := m[prop]; !exists {
		if customErrMsg != "" {
			return false, fmt.Errorf(customErrMsg, prop)
		}

		return false, fmt.Errorf("%q is not a valid option", prop)
	}

	return true, nil
}

package utils

import (
	"fmt"
)

func CheckValidMap[T comparable](m map[T]bool, prop T, customErrMsg string) (bool, error) {
	if _, exists := m[prop]; !exists {
		if customErrMsg != "" {
			return false, fmt.Errorf(customErrMsg, prop)
		}

		return false, fmt.Errorf("%v is not a valid option", prop)
	}

	return true, nil
}

package utils

import (
	"log"
	"os"
	"path/filepath"
)

func RelativeToPath(definedPaths ...string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Something went wrong with getting the current working directory")
	}

	parsedPath := filepath.Join(append([]string{cwd}, definedPaths...)...)

	return parsedPath
}

package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	bytes := make([]byte, 48)

	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	tokenToFeed := "om_" + hex.EncodeToString(bytes)

	// Define .env file path
	envFilePath := filepath.Join(filepath.Dir(os.Args[0]), "..", ".env")

	var envFileContent string

	if _, err := os.Stat(envFilePath); err == nil {
		data, err := os.ReadFile(envFilePath)
		if err != nil {
			fmt.Println("Error reading .env file:", err)
			return
		}

		envFileContent = string(data)
	} else {
		file, err := os.Create(envFilePath)
		if err != nil {
			fmt.Println("Error creating .env file:", err)
			return
		}

		defer file.Close()
	}

	if strings.Contains(envFileContent, "API_KEY=") {
		envFileContent = strings.ReplaceAll(envFileContent, `API_KEY=.*`, "API_KEY="+tokenToFeed)
	} else {
		envFileContent += "\nAPI_KEY=" + tokenToFeed
	}

	err = os.WriteFile(envFilePath, []byte(strings.TrimSpace(envFileContent)), 0644)
	if err != nil {
		fmt.Println("Error writing to .env file:", err)
		return
	}

	fmt.Printf("Token generated! This token has already populated from an .env file.\n\n%s\n", tokenToFeed)
}

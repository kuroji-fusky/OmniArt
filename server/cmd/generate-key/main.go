package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
)

func main() {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("For some reason, it couldn't generate a token:", err)
	}

	tokenToFeed := "om_" + hex.EncodeToString(bytes)

	envFilePath := utils.RelativeToPath("..", ".env")

	var envFileContent string

	if data, err := os.ReadFile(envFilePath); err == nil {
		envFileContent = string(data)
	} else if !os.IsNotExist(err) {
		log.Fatal("Error reading .env file:", err)
	}

	re := regexp.MustCompile(`(?m)^OM_SECRET_API_TOKEN=.*$`)

	if re.MatchString(envFileContent) {
		envFileContent = re.ReplaceAllString(envFileContent, "OM_SECRET_API_TOKEN="+tokenToFeed)
	} else {
		if len(envFileContent) > 0 {
			envFileContent += "\n"
		}
		envFileContent += "OM_SECRET_API_TOKEN=" + tokenToFeed
	}

	err = os.WriteFile(envFilePath, []byte(envFileContent), 0644)
	if err != nil {
		log.Fatal("Error writing to .env file:", err)
	}

	fmt.Printf("Token generated! This token has already populated from an .env file.\n\n%s\n", tokenToFeed)
}

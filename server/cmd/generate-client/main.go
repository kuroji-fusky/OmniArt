/*
A generator script that converts Go structs to TypeScript types
for the client to have a share of pizza and eat it too
*/
package main

import "github.com/kuroji-fusky/OmniArt/server/internal/utils"

func primitivesToTypescript(primitiveType string) string {
	switch primitiveType {
	case "string":
		return "string"

	case "int", "float32":
		return "number"

	case "bool":
		return "boolean"

	default:
		return "unknown"
	}
}

func main() {
	clientDir := utils.RelativeToPath("..", "client", "src")
}

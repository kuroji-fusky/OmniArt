package main

func convertToTs(primitiveType string) string {
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

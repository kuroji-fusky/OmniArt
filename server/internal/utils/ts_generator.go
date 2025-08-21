package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type GTSOptions struct {
	UnsafeAny bool // If you enable this, I will choke you
}

func GenerateTSType(v interface{}, opts *GTSOptions) string {
	t := reflect.TypeOf(v)
	out := "interface " + t.Name() + " {\n"

	typeMapping := map[string]string{
		"int":      "number",
		"string":   "string",
		"bool":     "boolean",
		"[]string": "string[]", // TODO refactor this to dynamically detect if a struct prop has `[]` on it
	}

	for i := range t.NumField() {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			continue
		}

		jsonParts := strings.Split(jsonTag, ",")

		jsonName := jsonParts[0]
		isOptional := len(jsonParts) > 1 && jsonParts[1] == "omitempty"

		// Wrap the prop with quotes if it has dashes
		if strings.Contains(jsonName, "-") {
			jsonName = `"` + jsonName + `"`
		}

		tsType := field.Tag.Get("ts_type")

		inferredType, exists := typeMapping[field.Type.Name()]
		if !exists {
			if opts != nil && opts.UnsafeAny {
				inferredType = "any"
			}

			inferredType = "unknown"
		}

		if tsType != "" {
			tsType = strings.ReplaceAll(tsType, "%t", inferredType)
		} else {
			tsType = inferredType
		}

		readonlyPrefix := "readonly "
		hasReadonlyPrefix := strings.Contains(tsType, "readonly")

		if !hasReadonlyPrefix {
			readonlyPrefix = ""
			tsType = strings.ReplaceAll(tsType, "readonly", "")
		}

		optionalSuffix := "?"
		if !isOptional {
			optionalSuffix = ""
		}

		out += fmt.Sprintf("  %s%s%s: %s;\n", readonlyPrefix, jsonName, optionalSuffix, tsType)
	}

	out += "}"
	return out
}

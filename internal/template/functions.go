package template

import (
	"regexp"
	"strings"
)

var Functions = map[string]interface{}{
	"toCamelCase":  toCamelCase,
	"toPascalCase": toPascalCase,
	"toSnakeCase":  toSnakeCase,
	"toKebabCase":  toKebabCase,
}

func toCamelCase(input string) string {
	return capitalizeWords(input, 1)
}

func toPascalCase(input string) string {
	return capitalizeWords(input, 0)
}

func capitalizeWords(input string, startingFrom int) string {
	normalized := normalize(input)
	for i := startingFrom; i < len(normalized); i++ {
		normalized[i] = capitalize(normalized[i])
	}
	joined := strings.Join(normalized, "")

	return joined
}

func toSnakeCase(input string) string {
	return joinWords(input, "_")
}

func toKebabCase(input string) string {
	return joinWords(input, "-")
}

func joinWords(input string, delimiter string) string {
	normalized := normalize(input)
	joined := strings.Join(normalized, delimiter)

	return joined
}

var delimiters = regexp.MustCompile(`([a-z])([A-Z])|[\s_\-]+`)

func normalize(input string) []string {
	delimitered := delimiters.ReplaceAllString(input, "${1} ${2}")
	lowered := strings.ToLower(delimitered)
	trimmed := strings.TrimSpace(lowered)
	split := strings.Fields(trimmed)

	return split
}

func capitalize(input string) string {
	if input == "" {
		return ""
	}
	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}

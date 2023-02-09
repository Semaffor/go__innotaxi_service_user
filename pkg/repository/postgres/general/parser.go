package general

import (
	"fmt"
	"strings"
)

// title?
func extractFieldsAndArgs(params map[string]interface{}) ([]string, []interface{}) {
	fields := make([]string, 0)
	args := make([]interface{}, 0)

	for field, value := range params {
		fields = append(fields, field)
		args = append(args, value)
	}

	return fields, args
}

func extractParamsWithDollar(params map[string]interface{}) (string, []interface{}) {
	fields, args := extractFieldsAndArgs(params)
	fieldsWithDollar := appendDollarToFields(fields)
	fieldsStr := strings.Join(fieldsWithDollar, ", ")

	return fieldsStr, args
}

func extractParams(params map[string]interface{}) (string, []interface{}) {
	fields, args := extractFieldsAndArgs(params)

	return strings.Join(fields, ", "), args
}

func appendDollarToFields(fields []string) []string {
	fieldsWithDollar := make([]string, 0)
	argId := 1

	for _, field := range fields {
		fieldsWithDollar = append(fieldsWithDollar, fmt.Sprintf("%s=$%d", field, argId))
		argId++
	}

	return fieldsWithDollar
}

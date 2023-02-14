package general

import (
	"fmt"
	"strings"
)

type paramsManipulator struct {
}

func newParamsManipulator() *paramsManipulator {
	return &paramsManipulator{}
}

// extractFieldsAndArgs divides input params on fields names and corresponding args.
func (m *paramsManipulator) extractFieldsAndArgs(params map[string]interface{}) ([]string, []interface{}) {
	fields := make([]string, 0)
	args := make([]interface{}, 0)

	for field, value := range params {
		fields = append(fields, field)
		args = append(args, value)
	}

	return fields, args
}

// extractFieldsAndArgs equals to extractFieldsAndArgs but fields appended with additional symbol.
func (m *paramsManipulator) extractParamsWithDollar(params map[string]interface{}) (string, []interface{}) {
	fields, args := m.extractFieldsAndArgs(params)
	fieldsWithDollar := m.appendDollarToFields(fields)
	fieldsStr := strings.Join(fieldsWithDollar, ", ")

	return fieldsStr, args
}

// extractParams just creates string from params with delimiter between them.
func (m *paramsManipulator) extractParams(params map[string]interface{}) (string, []interface{}) {
	fields, args := m.extractFieldsAndArgs(params)

	return strings.Join(fields, ", "), args
}

// appendDollarToFields add symbol '$1 ... $n' to the fields.
func (m *paramsManipulator) appendDollarToFields(fields []string) []string {
	fieldsWithDollar := make([]string, 0)
	argId := 1

	for _, field := range fields {
		fieldsWithDollar = append(fieldsWithDollar, fmt.Sprintf("%s=$%d", field, argId))
		argId++
	}

	return fieldsWithDollar
}

package builder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	table          string
	params         map[string]interface{}
	fields         []string
	args           []interface{}
	dollarSequence string
}

func NewQueryBuilder(table string, params map[string]interface{}) *QueryBuilder {
	return &QueryBuilder{
		table:  table,
		params: params,
		fields: make([]string, 0, len(params)),
		args:   make([]interface{}, 0, len(params)),
	}
}

// ExtractFieldsAndArgs divides input params on fields names and corresponding args.
func (qb *QueryBuilder) ExtractFieldsAndArgs() *QueryBuilder {
	for field, value := range qb.params {
		qb.fields = append(qb.fields, field)
		qb.args = append(qb.args, value)
	}

	return qb
}

// ExtractParamsWithDollar equals to extractFieldsAndArgs but fields appended with additional symbol.
func (qb *QueryBuilder) ExtractParamsWithDollar() *QueryBuilder {
	qb.AddDollarToFields()
	qb.separateFields()

	return qb
}

// separateFields just creates string from params with delimiter between them.
func (qb *QueryBuilder) separateFields() string {
	return strings.Join(qb.fields, ", ")
}

// AddDollarToFields add symbol '$1 ... $n' to the fields.
func (qb *QueryBuilder) AddDollarToFields() *QueryBuilder {
	fieldsWithDollar := make([]string, 0)
	argId := 1

	for _, field := range qb.fields {
		fieldsWithDollar = append(fieldsWithDollar, fmt.Sprintf("%s=$%d", field, argId))
		argId++
	}
	qb.fields = fieldsWithDollar

	return qb
}

// generateDollarSequence helps to create string in the following way: '$1, $2, ..., $n'
// for safe query injection.
func (qb *QueryBuilder) generateDollarSequence() *QueryBuilder {
	args := make([]string, 0)
	for i := 1; i <= len(qb.args); i++ {
		args = append(args, fmt.Sprintf("$%d", i))
	}
	qb.dollarSequence = strings.Join(args, ", ")

	return qb
}

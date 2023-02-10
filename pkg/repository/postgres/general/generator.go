package general

import (
	"fmt"
)

func GenerateSelectQuery(table string, params map[string]interface{}) (string, []interface{}) {
	fields, args := extractParamsWithDollar(params)

	if params == nil {
		return fmt.Sprintf("SELECT * FROM %s", table), nil
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s", table, fields), args
}

func GenerateUpdateQuery(table string, params map[string]interface{}, userId int) (string, []interface{}) {
	fields, args := extractParamsWithDollar(params)

	return fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id=%d", table, fields, userId), args
}

func GenerateInsertQuery(table string, params map[string]interface{}) (string, []interface{}) {
	fields, args := extractParams(params)
	dollars := generateDollarSequence(len(args))

	return fmt.Sprintf("INSERT INTO %s (%s) values (%v) returning id", table, fields, dollars), args
}

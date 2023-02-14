package general

import (
	"fmt"
)

// GenerateSelectQuery generate safe query with pointed params to get data in the table in database.
func GenerateSelectQuery(table string, params map[string]interface{}) (string, []interface{}) {
	manipulator := newParamsManipulator()
	fields, args := manipulator.extractParamsWithDollar(params)

	if params == nil {
		return fmt.Sprintf("SELECT * FROM %s", table), nil
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s", table, fields), args
}

// GenerateUpdateQuery generate safe update query with pointed params to update data in the table in database.
func GenerateUpdateQuery(table string, params map[string]interface{}, userId int) (string, []interface{}) {
	manipulator := newParamsManipulator()
	fields, args := manipulator.extractParamsWithDollar(params)

	return fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id=%d", table, fields, userId), args
}

// GenerateInsertQuery generate insert query with pointed params to update data in the table in database.
func GenerateInsertQuery(table string, params map[string]interface{}) (string, []interface{}) {
	manipulator := newParamsManipulator()
	fields, args := manipulator.extractParams(params)
	dollars := generateDollarSequence(len(args))

	return fmt.Sprintf("INSERT INTO %s (%s) values (%v) returning id", table, fields, dollars), args
}

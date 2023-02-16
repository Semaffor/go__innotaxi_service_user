package builder

import (
	"fmt"
)

// GenerateSelectQuery generate safe query with pointed params to get data in the table in database.
func (qb *QueryBuilder) GenerateSelectQuery() (string, []interface{}) {
	qb.ExtractParamsWithDollar()
	fieldsAsString := qb.separateFields()

	if qb.params == nil {
		return fmt.Sprintf("SELECT * FROM %s", qb.table), nil
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s", qb.table, fieldsAsString), qb.args
}

// GenerateUpdateQuery generate safe update query with pointed params to update data in the table in database.
func (qb *QueryBuilder) GenerateUpdateQuery(userId int) (string, []interface{}) {
	qb.ExtractParamsWithDollar()
	fieldsAsString := qb.separateFields()

	return fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id=%d", qb.table, fieldsAsString, userId), qb.args
}

// GenerateInsertQuery generate insert query with pointed params to update data in the table in database.
func (qb *QueryBuilder) GenerateInsertQuery() (string, []interface{}) {
	qb.generateDollarSequence()
	fieldsAsString := qb.separateFields()

	return fmt.Sprintf("INSERT INTO %s (%s) values (%v) returning id", qb.table, fieldsAsString, qb.dollarSequence),
		qb.args
}

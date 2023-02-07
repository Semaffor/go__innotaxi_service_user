package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

func executeQuery(db *sqlx.DB, query string, params []interface{}) error {
	var (
		res sql.Result
		err error
	)
	log.Printf("Executing query: %s with params: %s", query, params)

	if params == nil {
		res, err = db.Exec(query)
	} else {
		res, err = db.Exec(query, params...)
	}

	if err != nil {
		return errors.New(fmt.Sprintf("error when executing query: %s", query))
	}

	count, err := res.RowsAffected()
	if err != nil {
		return errors.New("error reading affected rows count")
	}

	if count == 0 {
		return errors.New(fmt.Sprintf("Affected rows count = 0"))
	}

	return nil
}

func generateSelectQuery(table string, params map[string]interface{}) (string, []interface{}) {
	fields, args := extractFieldsAndArgs(params)
	if params == nil {
		return fmt.Sprintf("SELECT * FROM %s", table), nil
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s", table, fields), args
}

func generateUpdateQuery(table string, params map[string]interface{}) (string, []interface{}) {
	fields, args := extractFieldsAndArgs(params)

	return fmt.Sprintf("UPDATE %s tl SET %s ", table, fields), args
}

func extractFieldsAndArgs(params map[string]interface{}) (string, []interface{}) {
	paramsStr := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	for field, value := range params {
		paramsStr = append(paramsStr, fmt.Sprintf("%s=$%d", field, argId))
		args = append(args, value)
		argId++
	}

	fields := strings.Join(paramsStr, ", ")

	return fields, args
}

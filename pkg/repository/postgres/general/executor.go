package general

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func ExecuteQuery(db *sqlx.DB, query string, params []interface{}) error {
	var (
		res sql.Result
		err error
	)
	log.Printf("Executing query: %s with params: %v", query, params)

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

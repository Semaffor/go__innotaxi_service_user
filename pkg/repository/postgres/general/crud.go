package general

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general/builder"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type ReturnType interface {
	model.User
}

type DaoI[T ReturnType] interface {
	Save(params map[string]interface{}) (int, error)
	Update(params map[string]interface{}, id int) error
	FindByFields(entities []T, params map[string]interface{}) ([]T, error)
	FindOneByFields(params map[string]interface{}) (*T, error)
}

type Dao[T ReturnType] struct {
	Db    *sqlx.DB
	Table string
}

// Save method generate insert query with params pointed in input map
// in the following way: key = field name, value = value to insert in db.
func (d *Dao[T]) Save(params map[string]interface{}) (int, error) {
	query, args :=
		builder.NewQueryBuilder(d.Table, params).
			ExtractFieldsAndArgs().
			GenerateDollarSequence().
			SeparateFields().
			GenerateInsertQuery()

	row := d.Db.QueryRowx(query, args...)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// FindByFields generic method which generate select query by pointed map values'.
func (d *Dao[T]) FindByFields(entities []T, params map[string]interface{}) ([]T, error) {
	var err error
	query, args := builder.NewQueryBuilder(d.Table, params).
		ExtractFieldsAndArgs().
		AddDollarToFields().
		SeparateFields().
		GenerateSelectQuery()

	if args != nil {
		err = d.Db.Select(&entities, query, args...)
	} else {
		err = d.Db.Select(&entities, query)
	}
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// Update generate query and update entity by id in database.
func (d *Dao[T]) Update(params map[string]interface{}, id int) error {
	query, args := builder.NewQueryBuilder(d.Table, params).
		ExtractFieldsAndArgs().
		AddDollarToFields().
		SeparateFields().
		GenerateUpdateQuery(id)

	err := ExecuteQuery(d.Db, query, args)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByFields performs the search of a single result by pointed fields,
// otherwise returning error.
func (d *Dao[T]) FindOneByFields(params map[string]interface{}) (*T, error) {
	entities, err := d.FindByFields([]T{}, params)
	if err != nil {
		return nil, err
	}
	if len(entities) > 1 {
		return nil, errors.New("more than 1 entity found")
	}
	if len(entities) == 0 {
		return &T{}, nil
	}

	return &entities[0], err
}

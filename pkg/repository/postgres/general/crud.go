package general

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

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

func (d *Dao[T]) Save(params map[string]interface{}) (int, error) {
	query, args := GenerateInsertQuery(d.Table, params)
	log.Print(query)
	row := d.Db.QueryRowx(query, args...)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (d *Dao[T]) FindByFields(entities []T, params map[string]interface{}) ([]T, error) {
	var err error
	query, args := GenerateSelectQuery(d.Table, params)

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

func (d *Dao[T]) Update(params map[string]interface{}, id int) error {
	query, args := GenerateUpdateQuery(d.Table, params, id)
	err := ExecuteQuery(d.Db, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao[T]) FindOneByFields(params map[string]interface{}) (*T, error) {
	entities, err := d.FindByFields([]T{}, params)
	if err != nil {
		return nil, err
	}

	if len(entities) > 1 {
		return nil, errors.New("more than 1 entity found")
	}

	return &entities[0], err
}

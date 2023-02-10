package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type DriverRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.Driver]
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{
		db:  db,
		dao: NewDriverDao(db),
	}
}

func (r *DriverRepository) Save(ctx context.Context, driver *model.Driver) (int, error) {
	params := map[string]interface{}{
		"user_id":      driver.UserId,
		"status":       model.FREE,
		"taxi_type_id": driver.TaxiTypeId,
	}

	return r.dao.Save(params)
}

func (r *DriverRepository) Update(ctx context.Context, driver *model.Driver) error {
	params := map[string]interface{}{}

	if driver.Status != 0 {
		params["status"] = driver.Status
	}

	return r.dao.Update(params, driver.Id)
}

func (r *DriverRepository) FindAll(ctx context.Context) ([]model.Driver, error) {
	return r.dao.FindByFields([]model.Driver{}, nil)
}

func (r *DriverRepository) FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Driver, error) {
	return r.dao.FindByFields([]model.Driver{}, params)
}

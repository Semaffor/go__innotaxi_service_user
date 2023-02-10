package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type TaxiTypeRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.TaxiType]
}

func NewTaxiTypeRepository(db *sqlx.DB) *TaxiTypeRepository {
	return &TaxiTypeRepository{
		db:  db,
		dao: NewTaxiTypeDao(db),
	}
}

func (r *TaxiTypeRepository) Save(ctx context.Context, taxiType *model.TaxiType) (int, error) {
	params := map[string]interface{}{
		"type":        taxiType.TaxiTypeTitle,
		"description": taxiType.Description,
	}

	return r.dao.Save(params)
}

func (r *TaxiTypeRepository) Update(ctx context.Context, role *model.TaxiType) error {
	params := map[string]interface{}{}

	if role.Description != "" {
		params["usr_id"] = role.Description
	}

	if role.TaxiTypeTitle != "" {
		params["role_id"] = role.TaxiTypeTitle
	}

	return r.dao.Update(params, role.Id)
}

func (r *TaxiTypeRepository) FindAll(ctx context.Context) ([]model.TaxiType, error) {
	return r.dao.FindByFields([]model.TaxiType{}, nil)
}

func (r *TaxiTypeRepository) FindByFields(
	ctx context.Context,
	params map[string]interface{},
) ([]model.TaxiType, error) {
	return r.dao.FindByFields([]model.TaxiType{}, params)
}

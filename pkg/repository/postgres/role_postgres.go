package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type RoleRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.Role]
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{
		db:  db,
		dao: NewRoleDao(db),
	}
}

func (r *RoleRepository) Save(ctx context.Context, role *model.Role) (int, error) {
	params := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
	}

	return r.dao.Save(params)
}

func (r *RoleRepository) Update(ctx context.Context, role *model.Role) error {
	params := map[string]interface{}{}

	if role.Name != "" {
		params["name"] = role.Name
	}

	if role.Description != "" {
		params["description"] = role.Description
	}

	return r.dao.Update(params, role.Id)
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]model.Role, error) {
	return r.dao.FindByFields([]model.Role{}, nil)
}

func (r *RoleRepository) FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Role, error) {
	return r.dao.FindByFields([]model.Role{}, params)
}

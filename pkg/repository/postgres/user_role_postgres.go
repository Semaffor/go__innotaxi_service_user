package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type UserRoleRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.UserRole]
}

func NewUserRoleRepository(db *sqlx.DB) *UserRoleRepository {
	return &UserRoleRepository{
		db:  db,
		dao: NewUserRoleDao(db),
	}
}

func (r *UserRoleRepository) Save(ctx context.Context, userRole *model.UserRole) (int, error) {
	params := map[string]interface{}{
		"usr_id":  userRole.UserId,
		"role_id": userRole.RoleId,
	}

	return r.dao.Save(params)
}

func (r *UserRoleRepository) Update(ctx context.Context, role *model.UserRole) error {
	params := map[string]interface{}{}

	if role.UserId != 0 {
		params["usr_id"] = role.UserId
	}

	if role.RoleId != 0 {
		params["role_id"] = role.RoleId
	}

	return r.dao.Update(params, role.Id)
}

func (r *UserRoleRepository) FindAll(ctx context.Context) ([]model.UserRole, error) {
	return r.dao.FindByFields([]model.UserRole{}, nil)
}

func (r *UserRoleRepository) FindByFields(
	ctx context.Context,
	params map[string]interface{},
) ([]model.UserRole, error) {
	return r.dao.FindByFields([]model.UserRole{}, params)
}

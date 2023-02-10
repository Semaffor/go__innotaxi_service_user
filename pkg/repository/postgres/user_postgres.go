package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/dao"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type UserRepository struct {
	db      *sqlx.DB
	userDao general.DaoI[model.User]
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db:      db,
		userDao: dao.NewUserDao(db, usersTable),
	}
}

func (r *UserRepository) Save(ctx context.Context, user *model.User) (int, error) {
	params := map[string]interface{}{
		"name":          user.Name,
		"phone_number":  user.PhoneNumber,
		"email":         user.Email,
		"password_hash": user.PasswordHash,
		"role":          model.USER,
	}

	return r.userDao.Save(params)
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	params := map[string]interface{}{}

	if user.Id != 0 {
		params["id"] = user.Id
	}

	if user.Name != "" {
		params["name"] = user.Name
	}

	if user.Username.Valid {
		params["username"] = user.Username
	}

	if user.PhoneNumber != "" {
		params["phone_number"] = user.PhoneNumber
	}

	if user.Email != "" {
		params["email"] = user.Email
	}

	if user.PasswordHash != "" {
		params["password_hash"] = user.PasswordHash
	}

	return r.userDao.Update(params, user.Id)
}

func (r *UserRepository) DeleteUserById(ctx context.Context, userId int) error {
	params := map[string]interface{}{
		"is_deleted": true,
	}

	return r.userDao.Update(params, userId)
}

func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	return r.userDao.FindByFields([]model.User{}, nil)
}

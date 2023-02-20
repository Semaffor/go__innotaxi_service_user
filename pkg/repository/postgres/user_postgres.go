package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type UserRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.User]
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db:  db,
		dao: NewUserDao(db),
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

	return r.dao.Save(params)
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	params := map[string]interface{}{}

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

	return r.dao.Update(params, user.Id)
}

func (r *UserRepository) DeleteUserById(ctx context.Context, userId int) error {
	params := map[string]interface{}{
		"is_deleted": true,
	}

	return r.dao.Update(params, userId)
}

func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	return r.dao.FindByFields([]model.User{}, nil)
}

func (r *UserRepository) FindByFields(ctx context.Context, params map[string]interface{}) ([]model.User, error) {
	return r.dao.FindByFields([]model.User{}, params)
}

func (r *UserRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	params := map[string]interface{}{
		"phone_number": phoneNumber,
	}

	return r.dao.FindOneByFields(params)
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	params := map[string]interface{}{
		"username": username,
	}

	return r.dao.FindOneByFields(params)
}

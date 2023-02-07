package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *model.User) (int, error) {
	query := fmt.Sprintf("insert into %s (name, phone_number, email, password_hash, role)"+
		" values ($1, $2, $3, $4, $5) returning id;", usersTable)
	row := r.db.QueryRowx(query, user.Name, user.PhoneNumber, user.Email, user.PasswordHash, model.USER)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
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

	query, args := generateUpdateQuery(usersTable, params)
	query = fmt.Sprintf("%s WHERE id=%d", query, user.Id)
	err := executeQuery(r.db, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUserById(ctx context.Context, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_deleted=true WHERE id=%d", usersTable, userId)

	err := executeQuery(r.db, query, nil)
	if err != nil {
		return err
	}

	return nil
}

// FindByFields probably make generic.
func (r *UserRepository) FindByFields(ctx context.Context, params map[string]interface{}) ([]model.User, error) {
	var err error
	query, args := generateSelectQuery(usersTable, params)

	var users []model.User
	if args != nil {
		err = r.db.Select(&users, query, args...)
	} else {
		err = r.db.Select(&users, query)
	}
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindAll probably make generic.
func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	return r.FindByFields(ctx, nil)
}

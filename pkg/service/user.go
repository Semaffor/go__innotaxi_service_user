package service

import (
	"context"

	"github.com/guregu/null"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/errbase"
	form "github.com/Semaffor/go__innotaxi_service_user/pkg/handler/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/hash"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
	pgModel "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type UserService struct {
	repo   repository.Users
	hasher hash.PasswordHasher
}

func NewUserService(repo repository.Users, hasher hash.PasswordHasher) *UserService {
	return &UserService{repo: repo, hasher: hasher}
}

func (u *UserService) Authenticate(ctx context.Context, credentials *form.UserLoginInput) (*pgModel.User, error) {
	var (
		user *pgModel.User
		err  error
	)

	if credentials.Username == "" {
		user, err = u.repo.FindByPhoneNumber(ctx, credentials.PhoneNumber)
	} else {
		user, err = u.repo.FindByUsername(ctx, credentials.Username)
	}
	if err != nil && !user.IsDeleted {
		return nil, errbase.InvalidCredentialsError("invalid username/phone/password")
	}

	hashedPassword, err := u.hasher.Hash(credentials.Password)
	if err != nil {
		return nil, errbase.DefaultError(err)
	}

	if user.PasswordHash != hashedPassword {
		return nil, errbase.InvalidCredentialsError("invalid username/phone/password")
	}

	return user, nil
}

func (u *UserService) Register(ctx context.Context, formUser *form.UserRegistrationInput) error {
	hashedPassword, err := u.hasher.Hash(formUser.Password)
	if err != nil {
		return errbase.DefaultError(err)
	}

	findByPhoneNumber, err := u.repo.FindByPhoneNumber(ctx, formUser.PhoneNumber)
	if err != nil {
		return err
	}
	if findByPhoneNumber.Id != 0 && !findByPhoneNumber.IsDeleted {
		return errbase.AlreadyExistsError("phone_number", formUser.PhoneNumber)
	}

	_, err = u.repo.Save(ctx, &pgModel.User{
		PhoneNumber:  formUser.PhoneNumber,
		Email:        formUser.Email,
		PasswordHash: hashedPassword,
		Role:         pgModel.USER,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(ctx context.Context, formUser *form.UserUpdateInput) error {
	var username null.String
	if formUser.Username != "" {
		username = null.StringFrom(formUser.Username)
	}

	pgUser := pgModel.User{
		Id:          formUser.Id,
		Name:        formUser.Name,
		Username:    username,
		PhoneNumber: formUser.PhoneNumber,
		Email:       formUser.Email,
	}

	return u.repo.Update(ctx, &pgUser)
}

func (u *UserService) DeleteUser(ctx context.Context, userId int) error {
	return u.repo.DeleteUserById(ctx, userId)
}

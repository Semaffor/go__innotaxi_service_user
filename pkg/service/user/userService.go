package user

import (
	"github.com/guregu/null"

	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

var fakeUserStruct = &model.User{
	Id:           1,
	Username:     null.StringFrom("Dima"),
	PhoneNumber:  "111",
	Email:        "dd@mail.ru",
	PasswordHash: "1",
}

type UserService struct {
	repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Authentication(credentials *modelJwt.UserCredentials) (model.User, error) {
	// find user in db
	return *fakeUserStruct, nil
}

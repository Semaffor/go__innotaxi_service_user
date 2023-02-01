package postgres

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

var fakeUserStruct = &domain.User{
	Id:           1,
	Username:     "Dima",
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

func (u *UserService) Authentication(credentials *domain.UserCredentials) (domain.User, error) {
	// find user in db
	return *fakeUserStruct, nil
}

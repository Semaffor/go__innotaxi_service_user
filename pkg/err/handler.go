package err

import (
	"fmt"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

// ServiceError What the best name for package and this struct?
type ServiceError struct {
	msg string
}

func (e *ServiceError) Error() string {
	return e.msg
}

func EntityNotFoundError(id int) ServiceError {
	return ServiceError{msg: fmt.Sprintf("Entity with id = %d not found", id)}
}

func InvalidCredentialsError() ServiceError {
	return ServiceError{msg: "Invalid username/password"}
}

func TokenExpiredError() ServiceError {
	return ServiceError{msg: "Token has expired"}
}

func DBConnectionError(dbName string, config config.ConfigDB) ServiceError {
	return ServiceError{msg: fmt.Sprintf("Couldn't connect to %s, config: %s", dbName, config)}
}

func UnsupportedDBError(msg string) ServiceError {
	return ServiceError{msg: msg}
}

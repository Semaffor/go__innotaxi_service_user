package err

import (
	"fmt"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

// Exception What the best name for package and this struct?
type Exception struct {
	msg string
}

func (e *Exception) Error() string {
	return e.msg
}

func EntityNotFoundError(id int) Exception {
	return Exception{msg: fmt.Sprintf("Entity with id = %d not found", id)}
}

func InvalidCredentialsError() Exception {
	return Exception{msg: fmt.Sprintf("Invalid username/password")}
}

func TokenExpiredError() Exception {
	return Exception{msg: fmt.Sprintf("Token has expired")}
}

func DBConnectionError(DBName string, config config.ConfigDB) Exception {
	return Exception{msg: fmt.Sprintf("Couldn't connect to %s, config: %s", DBName, config)}
}

func UnsupportedDBError(msg string) Exception {
	return Exception{msg: msg}
}

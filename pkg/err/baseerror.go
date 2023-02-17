package err

import (
	"fmt"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

const (
	Other          = 1  // Unclassified error. This value is not printed in the error message.
	Invalid        = 2  // Invalid operation for this type of item.
	IO             = 3  // External I/O error such as network failure.
	Exist          = 4  // Item already exists.
	NotExist       = 5  // Item does not exist.
	Private        = 6  // Information withheld.
	Internal       = 7  // Internal error or inconsistency.
	BrokenLink     = 8  // Link target does not exist.
	Database       = 9  // Error from database.
	Validation     = 10 // Input validation error.
	Unanticipated  = 11 // Unanticipated error.
	InvalidRequest = 12 // Invalid Request
	Token          = 13
)

type CodeError struct {
	ServiceCode int    `json:"code"`
	HttpCode    int    `json:"httpCode"`
	Msg         string `json:"msg"`
}

type CodeErrorResponse struct {
	ServiceCode int    `json:"code,omitempty"`
	FullCode    string `json:"fullCode,omitempty"`
	Msg         string `json:"msg,omitempty"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{ServiceCode: code, Msg: msg}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		FullCode: fmt.Sprintf("%d : %d", e.ServiceCode, e.HttpCode),
		Msg:      e.Msg,
	}
}

func EntityNotFoundError(id int) error {
	return NewCodeError(NotExist, fmt.Sprintf("Entity with id = %d not found", id))
}

func DefaultError(msg string) error {
	return NewCodeError(Other, msg)
}

func InvalidCredentialsError(msg string) error {
	return NewCodeError(Validation, "Incorrect phone/password.")
}

func TokenExpiredError() error {
	return NewCodeError(Token, "Token has expired.")
}

func DBConnectionError(dbName string, config config.DBConfig) error {
	return NewCodeError(Database, fmt.Sprintf("Couldn't connect to %s, config: %v", dbName, config))
}

func IOError(path string) error {
	return NewCodeError(IO, fmt.Sprintf("Couldn't read/write %s", path))
}

func AlreadyExistsError(field, value string) error {
	return NewCodeError(Exist, fmt.Sprintf("Entity with %s = %s already exists", field, value))
}

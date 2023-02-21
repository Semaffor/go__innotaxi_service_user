package errbase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler/model"
)

const (
	other          = 1  // Unclassified error. This value is not printed in the error message.
	invalid        = 2  // Invalid operation for this type of item.
	io             = 3  // External I/O error such as network failure.
	exist          = 4  // Item already exists.
	notExist       = 5  // Item does not exist.
	private        = 6  // Information withheld.
	internal       = 7  // Internal error or inconsistency.
	brokenLink     = 8  // Link target does not exist.
	database       = 9  // Error from database
	validation     = 10 // Input validation error.
	unanticipated  = 11 // Unanticipated error.
	invalidRequest = 12 // Invalid Request
	token          = 13 // Token error

	successful = "successful"
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

func NewCodeError(serviceCode int, httpCode int, msg string) error {
	return &CodeError{ServiceCode: serviceCode, HttpCode: httpCode, Msg: msg}
}

func NewErrorResponse(ctx *gin.Context, err error) {
	log.Printf("Error: %v", err)
	switch e := err.(type) {
	case *CodeError:
		ctx.AbortWithStatusJSON(e.HttpCode, NewJSONResponse(e.Data()))
	default:
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}

func NewJSONResponse(err *CodeErrorResponse) *model.JSONGeneralResponse {
	return &model.JSONGeneralResponse{ServiceCode: err.ServiceCode, Msg: err.Msg, Data: "-"}
}

func NewJSONSuccessResponse(data interface{}) *model.JSONGeneralResponse {
	localDataResponse := data
	if data == nil {
		localDataResponse = "-"
	}
	return &model.JSONGeneralResponse{ServiceCode: 0, Msg: successful, Data: localDataResponse}
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
	return NewCodeError(notExist, http.StatusNotFound, fmt.Sprintf("Entity with id = %d not found", id))
}

func KeyNotFoundError(key string) error {
	return NewCodeError(notExist, http.StatusNotFound, fmt.Sprintf("Refresh key = %s not found", key))
}

func DatabaseError(msg string) error {
	return NewCodeError(database, http.StatusInternalServerError, msg)
}

func DefaultError(err error) error {
	return NewCodeError(other, http.StatusInternalServerError, err.Error())
}

func InvalidCredentialsError(msg string) error {
	return NewCodeError(validation, http.StatusUnprocessableEntity, msg)
}

func TokenError(msg string) error {
	return NewCodeError(token, http.StatusUnauthorized, msg)
}

func DBConnectionError(dbName string, config config.DBConfig) error {
	return NewCodeError(database, http.StatusInternalServerError,
		fmt.Sprintf("Couldn't connect to %s, config: %v", dbName, config))
}

func IOError(path string) error {
	return NewCodeError(io, http.StatusInternalServerError, fmt.Sprintf("Couldn't read/write %s", path))
}

func AlreadyExistsError(field, value string) error {
	return NewCodeError(exist, http.StatusConflict, fmt.Sprintf("Entity with %s = %s already exists", field, value))
}

func InvalidInput(err error) error {
	return NewCodeError(invalidRequest, http.StatusInternalServerError,
		fmt.Sprintf("invalid input body: %s", err.Error()))
}

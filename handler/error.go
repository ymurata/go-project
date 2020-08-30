package handler

import (
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	echo "github.com/labstack/echo/v4"
)

// ErrorResponse ...
type ErrorResponse struct {
	Messages []string `json:"messages"`
}

var (
	// ErrorResponse500 ...
	ErrorResponse500 = &ErrorResponse{
		Messages: []string{"サーバー内部エラー"},
	}
)

const (
	validateErrorMessage = "%s is invalid value"
)

// HTTPErrorHandler ...
func HTTPErrorHandler(err error, c echo.Context) {
	s, e := NewErrorResponse(err)
	c.Logger().Error(err)
	if err := c.JSON(s, e); err != nil { // MEMO: c.JSON 内部でのエラーをハンドリング
		c.Logger().Error(err)
	}
}

// NewErrorResponse ...
func NewErrorResponse(err error) (int, *ErrorResponse) {
	// TODO: difine error type
	// switch err {
	// case xxxx:
	// 	return http.StatusBadRequest, getErrorResponse(err)

	switch e := err.(type) {
	case validator.ValidationErrors:
		return http.StatusBadRequest, getValidatorErrorResponse(e)
	}

	return http.StatusInternalServerError, ErrorResponse500
}

func getErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Messages: []string{},
	}
}

func getValidatorErrorResponse(errs validator.ValidationErrors) *ErrorResponse {
	es := make([]string, len(errs))
	for i, err := range errs {
		es[i] = fmt.Sprintf(validateErrorMessage, strcase.ToSnake(err.Field()))
	}
	return &ErrorResponse{
		Messages: es,
	}
}

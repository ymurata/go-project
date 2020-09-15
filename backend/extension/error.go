package extension

import (
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	echo "github.com/labstack/echo/v4"

	"go-project/interface/controller"
)

// ErrorResponse ...
type ErrorResponse struct {
	Messages []string `json:"messages"`
}

var (
	// ErrResponse404 ...
	ErrResponse404 = &ErrorResponse{
		Messages: []string{"存在しません"},
	}
	// ErrResponse500 ...
	ErrResponse500 = &ErrorResponse{
		Messages: []string{"サーバー内部エラー"},
	}
)

const (
	validateErrorMessage = "%s は不正な値です"
)

// HTTPErrorHandler ...
func HTTPErrorHandler(err error, c echo.Context) {
	s, e := NewErrorResponse(err)
	c.Logger().Error(err)
	if err := c.JSON(s, e); err != nil {
		c.Logger().Error(err)
	}
}

// NewErrorResponse ...
func NewErrorResponse(err error) (int, *ErrorResponse) {
	switch err {
	case controller.ErrBind:
		return http.StatusBadRequest, getErrorResponse(err)
	case echo.ErrNotFound:
		return http.StatusNotFound, ErrResponse404
	}

	switch e := err.(type) {
	case validator.ValidationErrors:
		return http.StatusBadRequest, getValidatorErrorResponse(e)
	}

	return http.StatusInternalServerError, ErrResponse500
}

func getErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Messages: []string{
			err.Error(),
		},
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

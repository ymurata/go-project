package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// ErrorResponse ...
type ErrorResponse struct {
	Messages []string `json:"messages"`
}

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
	// }
	return http.StatusInternalServerError, &ErrorResponse{Messages: []string{"サーバー内部エラー"}}
}

func getErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Messages: []string{},
	}
}

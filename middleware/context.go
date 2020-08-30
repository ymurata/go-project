package middleware

import (
	echo "github.com/labstack/echo/v4"

	"go-project/context"
)

// Context set custom context for api
func Context() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&context.Context{Context: c})
		}
	}
}

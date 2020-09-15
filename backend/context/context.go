package context

import (
	"errors"

	echo "github.com/labstack/echo/v4"
)

var (
	// ErrBind ...
	ErrBind = errors.New("不正なリクエストです")
)

// Context ...
type Context struct {
	echo.Context
}

// BindAndValidate ...
func (c *Context) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return ErrBind
	}
	err := c.Validate(i)
	if errors.Is(err, echo.ErrValidatorNotRegistered) {
		c.Logger().Error(err)
	}
	return err
}

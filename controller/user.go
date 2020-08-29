package controller

import (
	"go-project/request"
	"go-project/service"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// UserController ...
type UserController struct {
	service service.UserService
}

// NewUserController ...
func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// List ...
func (u *UserController) List(c echo.Context) error {
	users, err := u.service.List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Get ...
func (u *UserController) Get(c echo.Context) error {
	var data request.UserID
	if err := c.Bind(&data); err != nil {
		return err
	}

	user, err := u.service.Get(data.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Create ...
func (u *UserController) Create(c echo.Context) error {
	var data request.UserCreate
	if err := c.Bind(&data); err != nil {
		return err
	}

	user, err := u.service.Create(data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Update ...
func (u *UserController) Update(c echo.Context) error {
	var data request.UserUpdate
	if err := c.Bind(&data); err != nil {
		return err
	}

	user, err := u.service.Update(1, data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Delete ...
func (u *UserController) Delete(c echo.Context) error {
	var data request.UserID
	if err := c.Bind(&data); err != nil {
		return err
	}
	return u.service.Delete(data.ID)
}

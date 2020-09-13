package controller

import (
	"net/http"

	"go-project/context"
	"go-project/interface/parameter"
	"go-project/service"
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
func (u *UserController) List(c context.Context) error {
	users, err := u.service.List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Get ...
func (u *UserController) Get(c context.Context) error {
	var data parameter.UserID
	if err := c.BindAndValidate(&data); err != nil {
		return err
	}

	user, err := u.service.Get(data.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Create ...
func (u *UserController) Create(c context.Context) error {
	var data parameter.UserCreate
	if err := c.BindAndValidate(&data); err != nil {
		return err
	}

	user, err := u.service.Create(data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Update ...
func (u *UserController) Update(c context.Context) error {
	var data parameter.UserUpdate
	if err := c.BindAndValidate(&data); err != nil {
		return err
	}

	user, err := u.service.Update(data.ID, data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Delete ...
func (u *UserController) Delete(c context.Context) error {
	var data parameter.UserID
	if err := c.BindAndValidate(&data); err != nil {
		return err
	}
	return u.service.Delete(data.ID)
}

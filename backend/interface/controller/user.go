package controller

import (
	"errors"
	"net/http"

	"go-project/context"
	"go-project/domain/model"
	"go-project/interface/parameter"
	"go-project/service"
)

type (
	// UserResponse ...
	UserResponse struct {
		User *model.User `json:"user"`
	}

	// UsersResponse ...
	UsersResponse struct {
		Users []*model.User `json:"users"`
	}

	// UserController ...
	UserController struct {
		service service.UserService
	}
)

// NewUserController ...
func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// List ...
func (u *UserController) List(ctx context.Context) error {
	users, err := u.service.List(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, UsersResponse{Users: users})
}

// Get ...
func (u *UserController) Get(ctx context.Context) error {
	var data parameter.UserID
	if err := ctx.BindAndValidate(&data); err != nil {
		if errors.Is(err, context.ErrBind) {
			return ErrBind
		}
		return err
	}

	user, err := u.service.Get(ctx, data)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, UserResponse{User: user})
}

// Create ...
func (u *UserController) Create(ctx context.Context) error {
	var data parameter.UserCreate
	if err := ctx.BindAndValidate(&data); err != nil {
		if errors.Is(err, context.ErrBind) {
			return ErrBind
		}
		return err
	}

	user, err := u.service.Create(ctx, data)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, UserResponse{User: user})
}

// Update ...
func (u *UserController) Update(ctx context.Context) error {
	var data parameter.UserUpdate
	if err := ctx.BindAndValidate(&data); err != nil {
		if errors.Is(err, context.ErrBind) {
			return ErrBind
		}
		return err
	}

	user, err := u.service.Update(ctx, data)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, UserResponse{User: user})
}

// Delete ...
func (u *UserController) Delete(ctx context.Context) error {
	var data parameter.UserID
	if err := ctx.BindAndValidate(&data); err != nil {
		if errors.Is(err, context.ErrBind) {
			return ErrBind
		}
		return err
	}
	if err := u.service.Delete(ctx, data); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, UserResponse{User: nil})
}

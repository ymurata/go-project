package service

import (
	"errors"

	"gorm.io/gorm"

	"go-project/context"
	"go-project/domain/model"
	"go-project/domain/repository"
	"go-project/infrastructure/database"
	"go-project/interface/parameter"
)

type (
	// UserService ...
	UserService interface {
		List(ctx context.Context) ([]*model.User, error)
		Get(ctx context.Context, data parameter.UserID) (*model.User, error)
		Create(ctx context.Context, data parameter.UserCreate) (*model.User, error)
		Update(ctx context.Context, data parameter.UserUpdate) (*model.User, error)
		Delete(ctx context.Context, data parameter.UserID) error
	}
	// UserServiceImpl ...
	UserServiceImpl struct {
		db   database.Handler
		repo repository.UserRepository
	}
)

// NewUserServiceImpl ...
func NewUserServiceImpl(db database.Handler, repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		db:   db,
		repo: repo,
	}
}

// List ...
func (u *UserServiceImpl) List(ctx context.Context) ([]*model.User, error) {
	return u.repo.List(u.db.Get())
}

// Get ...
func (u *UserServiceImpl) Get(ctx context.Context, data parameter.UserID) (*model.User, error) {
	user, err := u.repo.Get(u.db.Get(), data)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, nil
}

// Create ...
func (u *UserServiceImpl) Create(ctx context.Context, data parameter.UserCreate) (*model.User, error) {
	return u.repo.Create(u.db.Get(), data)
}

// Update ...
func (u *UserServiceImpl) Update(ctx context.Context, data parameter.UserUpdate) (*model.User, error) {
	user, err := u.repo.Update(u.db.Get(), data)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, nil
}

// Delete ...
func (u *UserServiceImpl) Delete(ctx context.Context, data parameter.UserID) error {
	err := u.repo.Delete(u.db.Get(), data)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

var _ UserService = (*UserServiceImpl)(nil)

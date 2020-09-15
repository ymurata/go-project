package service

import (
	"go-project/domain/model"
	"go-project/domain/repository"
	"go-project/extension"
	"go-project/infrastructure/database"
	"go-project/interface/parameter"
)

type (
	// UserService ...
	UserService interface {
		List() ([]*model.User, error)
		Get(id extension.HashID64) (*model.User, error)
		Create(data parameter.UserCreate) (*model.User, error)
		Update(id extension.HashID64, data parameter.UserUpdate) (*model.User, error)
		Delete(id extension.HashID64) error
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
func (u *UserServiceImpl) List() ([]*model.User, error) {
	return u.repo.List(u.db.Get())
}

// Get ...
func (u *UserServiceImpl) Get(id extension.HashID64) (*model.User, error) {
	return u.repo.Get(u.db.Get(), id)
}

// Create ...
func (u *UserServiceImpl) Create(data parameter.UserCreate) (*model.User, error) {
	return u.repo.Create(u.db.Get(), data)
}

// Update ...
func (u *UserServiceImpl) Update(id extension.HashID64, data parameter.UserUpdate) (*model.User, error) {
	return u.repo.Update(u.db.Get(), id, data)
}

// Delete ...
func (u *UserServiceImpl) Delete(id extension.HashID64) error {
	return u.repo.Delete(u.db.Get(), id)
}

var _ UserService = (*UserServiceImpl)(nil)

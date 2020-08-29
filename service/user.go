package service

import (
	"go-project/domain/model"
	"go-project/domain/repository"
	"go-project/request"
)

type (
	// UserService ...
	UserService interface {
		List() ([]*model.User, error)
		Get(id int64) (*model.User, error)
		Create(data request.UserCreate) (*model.User, error)
		Update(id int64, data request.UserUpdate) (*model.User, error)
		Delete(id int64) error
	}
	// UserServiceImpl ...
	UserServiceImpl struct {
		repo repository.UserRepository
	}
)

// NewUserServiceImpl ...
func NewUserServiceImpl(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

// List ...
func (u *UserServiceImpl) List() ([]*model.User, error) {
	return u.repo.List()
}

// Get ...
func (u *UserServiceImpl) Get(id int64) (*model.User, error) {
	return u.repo.Get(id)
}

// Create ...
func (u *UserServiceImpl) Create(data request.UserCreate) (*model.User, error) {
	return u.repo.Create(data)
}

// Update ...
func (u *UserServiceImpl) Update(id int64, data request.UserUpdate) (*model.User, error) {
	return u.repo.Update(id, data)
}

// Delete ...
func (u *UserServiceImpl) Delete(id int64) error {
	return u.repo.Delete(id)
}

var _ UserService = (*UserServiceImpl)(nil)

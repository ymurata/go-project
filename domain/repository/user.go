package repository

import (
	"gorm.io/gorm"

	"go-project/domain/model"
	"go-project/infrastructure/database"
	"go-project/request"
)

type (
	// UserRepository ...
	UserRepository interface {
		List() ([]*model.User, error)
		Get(id int64) (*model.User, error)
		Create(data request.UserCreate) (*model.User, error)
		Update(id int64, data request.UserUpdate) (*model.User, error)
		Delete(id int64) error
	}
	// UserRepositoryImpl ...
	UserRepositoryImpl struct {
		db *gorm.DB
	}
)

// NewUserRepositoryImpl ...
func NewUserRepositoryImpl(db *database.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db.Get(),
	}
}

// List ...
func (u *UserRepositoryImpl) List() ([]*model.User, error) {
	var users []*model.User
	return users, nil
}

// Get ...
func (u *UserRepositoryImpl) Get(id int64) (*model.User, error) {
	return nil, nil
}

// Create ...
func (u *UserRepositoryImpl) Create(data request.UserCreate) (*model.User, error) {
	return nil, nil
}

// Update ...
func (u *UserRepositoryImpl) Update(id int64, data request.UserUpdate) (*model.User, error) {
	return nil, nil
}

// Delete ...
func (u *UserRepositoryImpl) Delete(id int64) error {
	return nil
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

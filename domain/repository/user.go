package repository

import (
	"errors"

	"gorm.io/gorm"

	"go-project/domain/model"
	"go-project/extension"
	"go-project/infrastructure/database"
	"go-project/interface/parameter"
)

type (
	// UserRepository ...
	UserRepository interface {
		List() ([]*model.User, error)
		Get(id extension.HashID64) (*model.User, error)
		Create(data parameter.UserCreate) (*model.User, error)
		Update(id extension.HashID64, data parameter.UserUpdate) (*model.User, error)
		Delete(id extension.HashID64) error
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
	if err := u.db.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*model.User{}, nil
		}
		return []*model.User{}, err
	}
	return users, nil
}

// Get ...
func (u *UserRepositoryImpl) Get(id extension.HashID64) (*model.User, error) {
	return u.findByID(id)
}

// Create ...
func (u *UserRepositoryImpl) Create(data parameter.UserCreate) (*model.User, error) {
	user := model.User{
		Name:  data.Name,
		Email: data.Email,
	}
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update ...
func (u *UserRepositoryImpl) Update(id extension.HashID64, data parameter.UserUpdate) (*model.User, error) {
	var user model.User
	ud := model.User{Name: data.Name}
	if err := u.db.Model(&user).Where("id = ?", id).Updates(ud).Error; err != nil {
		return &user, err
	}
	return u.findByID(id)
}

// Delete ...
func (u *UserRepositoryImpl) Delete(id extension.HashID64) error {
	// TODO: add delete flg
	return nil
}

func (u *UserRepositoryImpl) findByID(id extension.HashID64) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.User{}, nil
		}
		return nil, err
	}
	return &user, nil
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

package repository

import (
	"errors"

	"gorm.io/gorm"

	"go-project/domain/model"
	"go-project/extension/types"
	"go-project/interface/parameter"
)

type (
	// UserRepository ...
	UserRepository interface {
		List(db *gorm.DB) ([]*model.User, error)
		Get(db *gorm.DB, data parameter.UserID) (*model.User, error)
		Create(db *gorm.DB, data parameter.UserCreate) (*model.User, error)
		Update(db *gorm.DB, data parameter.UserUpdate) (*model.User, error)
		Delete(db *gorm.DB, data parameter.UserID) error
	}

	// UserRepositoryImpl ...
	UserRepositoryImpl struct{}
)

// NewUserRepositoryImpl ...
func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

// List ...
func (u *UserRepositoryImpl) List(db *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	if err := db.Scopes(filterDeleted).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*model.User{}, nil
		}
		return []*model.User{}, err
	}
	return users, nil
}

// Get ...
func (u *UserRepositoryImpl) Get(db *gorm.DB, data parameter.UserID) (*model.User, error) {
	return u.findByID(db, data.ID)
}

// Create ...
func (u *UserRepositoryImpl) Create(db *gorm.DB, data parameter.UserCreate) (*model.User, error) {
	user := model.User{
		Name:  data.Name,
		Email: data.Email,
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update ...
func (u *UserRepositoryImpl) Update(db *gorm.DB, data parameter.UserUpdate) (*model.User, error) {
	var user model.User
	ud := model.User{Name: data.Name}
	if err := db.Model(&user).Where("id = ?", data.ID).Updates(ud).Error; err != nil {
		return &user, err
	}
	return u.findByID(db, data.ID)
}

// Delete ...
func (u *UserRepositoryImpl) Delete(db *gorm.DB, data parameter.UserID) error {
	user, err := u.findByID(db, data.ID)
	if err != nil {
		return err
	}
	return db.Model(user).Update("deleted", true).Error
}

func (u *UserRepositoryImpl) findByID(db *gorm.DB, id types.HashID64) (*model.User, error) {
	var user model.User
	if err := db.Scopes(filterDeleted).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

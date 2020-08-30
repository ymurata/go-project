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
	if err := u.db.Find(&users).Error; err != nil {
		return []*model.User{}, err
	}
	return users, nil
}

// Get ...
func (u *UserRepositoryImpl) Get(id int64) (*model.User, error) {
	return u.findByID(id)
}

// Create ...
func (u *UserRepositoryImpl) Create(data request.UserCreate) (*model.User, error) {
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
func (u *UserRepositoryImpl) Update(id int64, data request.UserUpdate) (*model.User, error) {
	var user model.User
	ud := model.User{Name: data.Name}
	if err := u.db.Model(&user).Where("id = ?", id).Updates(ud).Error; err != nil {
		return &user, err
	}
	return u.findByID(id)
}

// Delete ...
func (u *UserRepositoryImpl) Delete(id int64) error {
	// TODO: add delete flg
	return nil
}

func (u *UserRepositoryImpl) findByID(id int64) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

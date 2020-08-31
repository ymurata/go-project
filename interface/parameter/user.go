package parameter

import "go-project/extension"

// UserID ...
type UserID struct {
	ID extension.HashID64 `param:"id" validate:"required"`
}

// UserCreate ...
type UserCreate struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// UserUpdate ...
type UserUpdate struct {
	ID   extension.HashID64 `param:"id" validate:"required"`
	Name string             `json:"name" validate:"required"`
}

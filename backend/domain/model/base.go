package model

import (
	"time"

	"go-project/extension"
)

// Base ...
type Base struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BaseWithHashID ...
type BaseWithHashID struct {
	ID        extension.HashID64 `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

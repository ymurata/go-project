package model

import (
	"time"

	"go-project/extension/types"
)

// Base ...
type Base struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
}

// BaseWithHashID ...
type BaseWithHashID struct {
	ID        types.HashID64 `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"updated_at"`
	Deleted   bool           `json:"deleted"`
}

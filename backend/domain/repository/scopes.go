package repository

import "gorm.io/gorm"

func filterDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted", false)
}

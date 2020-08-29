package database

import (
	"database/sql"

	// pq package
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	db *gorm.DB
}

// New ...
func New(databaseURL string) (*DB, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{Conn: conn}),
		&gorm.Config{},
	)
	return &DB{db: db}, err
}

// Get ...
func (d *DB) Get() *gorm.DB {
	return d.db
}

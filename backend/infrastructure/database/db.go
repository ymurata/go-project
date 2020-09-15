package database

import (
	"database/sql"

	// pq package
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	// Handler ...
	Handler interface {
		Get() *gorm.DB
		Transaction(process func(tx *gorm.DB) (interface{}, error)) (interface{}, error)
	}

	// DB ...
	DB struct {
		db *gorm.DB
	}
)

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

// Transaction ...
func (d *DB) Transaction(process func(tx *gorm.DB) (interface{}, error)) (interface{}, error) {
	var result interface{}
	err := d.db.Transaction(func(tx *gorm.DB) error {
		data, err := process(tx)
		if err != nil {
			return err
		}
		result = data
		return nil
	})
	return result, err
}

var _ Handler = (*DB)(nil)

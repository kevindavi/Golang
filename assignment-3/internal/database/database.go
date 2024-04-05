package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:anjing123@localhost/orders_by?sslmode=disable")

	if err != nil {
		return nil, err
	}
	return db, nil
}

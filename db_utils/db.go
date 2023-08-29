package db_utils

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "postgresql://postgres:postgres@localhost:5432/dotfyle?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

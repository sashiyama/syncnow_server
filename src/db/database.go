package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func New() *sql.DB {
	connStr := os.Getenv("POSTGRESQL_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

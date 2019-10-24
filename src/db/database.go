package db

import (
	"database/sql"
	"log"
)

type Database struct {
	DriverName     string
	DataSourceName string
}

func (d *Database) New() *sql.DB {
	db, err := sql.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

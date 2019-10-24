package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

type Postgres struct {
	DriverName     string
	DataSourceName string
}

func NewPostgres() *sql.DB {
	p := Postgres{DriverName: "postgres", DataSourceName: os.Getenv("POSTGRESQL_URL")}
	d := Database{DriverName: p.DriverName, DataSourceName: p.DataSourceName}
	return d.New()
}

package util

import (
	"github.com/sashiyama/syncnow_server/db"
	"log"
)

func TruncateAllTables() {
	d := db.NewPostgres()
	rows, err := d.Query("SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
	if err != nil {
		log.Fatal("show tables error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			log.Fatal("show table error:", err)
		}
		_, err = d.Exec("TRUNCATE " + tableName + " CASCADE")
		if err != nil {
			log.Fatal("truncate table error:", err)
		}
	}
}

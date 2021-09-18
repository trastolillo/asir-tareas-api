package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	var ConnectionString = ""
	if IsLocal {
		ConnectionString = LocalConnectionString
	} else {
		ConnectionString = RemoteConnectionString
	}
	log.Printf("ConnectionString: %s", ConnectionString)
	return sql.Open("mysql", ConnectionString)
}

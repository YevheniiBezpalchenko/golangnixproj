package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() {

	db, err := sql.Open(
		"mysql",
		"user:12345@tcp(127.0.0.1:3306)/shop",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

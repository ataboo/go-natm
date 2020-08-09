package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewSqlDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ataboo/go-natm/pkg/common"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Comment prevents lint
	_ "github.com/lib/pq"

	// Comment prevents lint
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewSqlDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(filepath.Join("file://", common.RootFilePath, "migrations"), "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == nil {
		fmt.Println("Migrated DB")
	}

	if err == migrate.ErrNoChange {
		return nil
	}

	return err
}

package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/ataboo/go-natm/pkg/common"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func NewSqlTestDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv(common.EnvTestDbConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MustMigrateTestDB() *sql.DB {
	realDB := NewSqlDB()
	if _, err := realDB.Exec("DROP DATABASE IF EXISTS gonatm_test"); err != nil {
		log.Fatal(err)
	}
	if _, err := realDB.Exec("CREATE DATABASE gonatm_test"); err != nil {
		log.Fatal(err)
	}
	realDB.Close()

	testDB := NewSqlTestDB()
	driver, err := postgres.WithInstance(testDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(filepath.Join("file://", common.RootFilePath, "migrations"), "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

	return testDB
}

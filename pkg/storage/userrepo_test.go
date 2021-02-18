package storage

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ataboo/go-natm/pkg/database"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var test_db *sql.DB
var test_ctx context.Context
var test_tx *sql.Tx

func TestMain(m *testing.M) {
	repo_fixturesetup()
	code := m.Run()
	repo_fixtureteardown()
	os.Exit(code)
}

func repo_fixturesetup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	test_db = database.MustMigrateTestDB()
	test_ctx = context.Background()
}

func repo_fixtureteardown() {
	test_db.Close()
}

func repo_testsetup(t *testing.T) {
	newTx, err := test_db.BeginTx(test_ctx, &sql.TxOptions{})
	if err != nil {
		t.Error(err)
	}

	test_tx = newTx
}

func repo_testteardown() {
	test_tx.Rollback()
}

func TestUserRepoBasics(t *testing.T) {
	repo_testsetup(t)
	defer repo_testteardown()

	user := &models.User{
		Email:  "test@user.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Test User",
	}

	repo := NewUserRepository(test_db)
	err := repo.Create(user)
	if err != nil {
		t.Error(err)
	}

	readUser, err := repo.Find(user.ID)
	if err != nil {
		t.Error(err)
	}

	if readUser.Name != user.Name || readUser.ID != user.ID {
		t.Error("expected name to match")
	}

	emailUser := repo.FindByEmail("test@user.com")
	if emailUser == nil || emailUser.Name != user.Name {
		t.Error("failed to find user by email")
	}

	emailUser.Name = "changed name"
	err = repo.Update(emailUser)
	if err != nil {
		t.Error(err)
	}

	changedUser, err := repo.Find(user.ID)
	if err != nil {
		t.Error(err)
	}

	if changedUser.Name != "changed name" {
		t.Error("expected name to be changed")
	}
}

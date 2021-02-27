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

var testDb *sql.DB
var testCtx context.Context
var testTx *sql.Tx

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

	testDb = database.MustMigrateTestDB()
	testCtx = context.Background()
}

func repo_fixtureteardown() {
	testDb.Close()
}

func repo_testsetup(t *testing.T) {
	newTx, err := testDb.BeginTx(testCtx, &sql.TxOptions{})
	if err != nil {
		t.Error(err)
	}

	testTx = newTx
}

func repo_testteardown() {
	testTx.Rollback()
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

	repo := NewUserRepository(testDb)
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

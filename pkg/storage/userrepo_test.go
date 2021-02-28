package storage

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/ataboo/go-natm/pkg/database"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var testDb *sql.DB
var testCtx context.Context
var testTx *sql.Tx
var testDbLock sync.Mutex

func TestMain(m *testing.M) {
	repo_fixturesetup()
	code := m.Run()
	repo_fixtureteardown()
	os.Exit(code)
}

func repo_fixturesetup() {
	testDbLock = sync.Mutex{}
	testCtx = context.Background()
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	testDb = database.MustMigrateTestDB()
}

func repo_fixtureteardown() {
	testDb.Close()
}

func repo_testsetup(t *testing.T) {
	testDbLock.Lock()

	// _, file, no, _ := runtime.Caller(1)
	// users, _ := models.Users().All(testCtx, testDb)
	// fmt.Printf("%s:%d: %d users before tx\n", file, no, len(users))

	newTx, err := testDb.BeginTx(testCtx, nil)
	if err != nil {
		t.Error(err)
	}

	testTx = newTx
}

func repo_testteardown() {
	if err := testTx.Rollback(); err != nil {
		log.Fatal(err)
	}

	// users, _ := models.Users().All(testCtx, testDb)
	// fmt.Printf("%d users after rollback\n", len(users))

	defer testDbLock.Unlock()
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

	repo := NewUserRepository(testCtx, testTx)
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

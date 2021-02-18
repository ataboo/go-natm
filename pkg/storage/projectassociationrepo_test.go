package storage

import (
	"context"
	"testing"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestFoo(t *testing.T) {
	repo_testsetup(t)
	repo_testteardown()

	repo := NewProjectAssociationRepository(test_db)

	user := models.User{
		Email:  "test@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Test User",
	}

	project := models.Project{}

	if err := user.Insert(context.Background(), test_db, boil.Infer()); err != nil {
		t.Error(err)
	}

	err := repo.Create(&data.ProjectAssociationCreate{}, user.ID)

	t.Error(err)
}

package storage

import (
	"testing"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestCreateProjectAndListAssociated(t *testing.T) {
	repo_testsetup(t)
	defer repo_testteardown()

	owner := &models.User{
		Email:  "user1@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Project Owner",
	}
	if err := owner.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	secondUser := &models.User{
		Email:  "user2@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Second User",
	}
	if err := secondUser.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	thirdUser := &models.User{
		Email:  "user3@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Third User",
	}
	if err := thirdUser.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	projectRepo := NewProjectRepository(testTx)

	if err := projectRepo.Create(&data.ProjectCreate{}, uuid.NewString()); err == nil {
		t.Error("expected create fail with invalid user id")
	}

	if err := projectRepo.Create(&data.ProjectCreate{}, owner.ID); err == nil {
		t.Error("expected create fail for no string data")
	}

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "Zeta First Project",
		Abbreviation: "PRJ1",
		Description:  "First test project.",
	}, owner.ID); err != nil {
		t.Error(err)
	}

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "Alpha Second Project",
		Abbreviation: "PRJ2",
		Description:  "Second test project.",
	}, secondUser.ID); err != nil {
		t.Error(err)
	}

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "Third Project",
		Abbreviation: "PRJ3",
		Description:  "Third test project",
	}, thirdUser.ID); err != nil {
		t.Error(err)
	}

	firstProject, err := models.Projects(qm.Where("name = ?", "Zeta First Project")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}
	secondProject, err := models.Projects(qm.Where("name = ?", "Alpha Second Project")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}

	associations, err := projectRepo.ListAssociated(owner.ID)
	if err != nil {
		t.Error(err)
	}

	if len(associations) != 1 || associations[0].ProjectID != firstProject.ID {
		t.Error("expected first project association")
	}

	writerAssoc := models.ProjectAssociation{
		ID:          uuid.NewString(),
		ProjectID:   secondProject.ID,
		Email:       owner.Email,
		Association: models.AssociationsEnumWriter,
	}
	if err := writerAssoc.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	associations, err = projectRepo.ListAssociated(owner.ID)
	if err != nil {
		t.Error(err)
	}
	if len(associations) != 2 || associations[0].ProjectID != firstProject.ID || associations[1].ProjectID != secondProject.ID {
		t.Error("expected first project association")
	}

	_, err = projectRepo.ListAssociated(uuid.NewString())
	if err == nil {
		t.Error("expected invalid user id")
	}
}

func TestProjectFind(t *testing.T) {
	repo_testsetup(t)
	defer repo_testteardown()

	owner := &models.User{
		Email:  "zuser1@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Project Owner",
	}
	if err := owner.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	secondUser := &models.User{
		Email:  "auser2@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Second User",
	}
	if err := secondUser.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	thirdUser := &models.User{
		Email:  "user3@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Third User",
	}
	if err := thirdUser.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	projectRepo := NewProjectRepository(testTx)

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "First Project",
		Abbreviation: "PRJ1",
		Description:  "First test project.",
	}, owner.ID); err != nil {
		t.Error(err)
	}

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "Second Project",
		Abbreviation: "PRJ2",
		Description:  "Second test project.",
	}, secondUser.ID); err != nil {
		t.Error(err)
	}

	if err := projectRepo.Create(&data.ProjectCreate{
		Name:         "Third Project",
		Abbreviation: "PRJ3",
		Description:  "Third test project",
	}, thirdUser.ID); err != nil {
		t.Error(err)
	}

	firstProject, err := models.Projects(qm.Where("name = ?", "First Project")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}
	secondProject, err := models.Projects(qm.Where("name = ?", "Second Project")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}
	thirdProject, err := models.Projects(qm.Where("name = ?", "Third Project")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}

	writerAssoc := models.ProjectAssociation{
		ID:          uuid.NewString(),
		ProjectID:   secondProject.ID,
		Email:       owner.Email,
		Association: models.AssociationsEnumWriter,
	}
	if err := writerAssoc.Insert(testCtx, testTx, boil.Infer()); err != nil {
		t.Error(err)
	}

	_, err = projectRepo.Find(firstProject.ID, uuid.NewString())
	if err == nil {
		t.Error("expected failed to find with invalid user id")
	}

	association, err := projectRepo.Find(firstProject.ID, owner.ID)
	if err != nil {
		t.Error(err)
	}

	if association.ProjectID != firstProject.ID {
		t.Error("unnexpected project id")
	}

	association, err = projectRepo.Find(secondProject.ID, owner.ID)
	if err != nil {
		t.Error(err)
	}

	if association.ProjectID != secondProject.ID {
		t.Error("unnexpected project id")
	}

	association, err = projectRepo.Find(thirdProject.ID, owner.ID)
	if err == nil || association != nil {
		t.Error("expected not result for 3rd project")
	}
}

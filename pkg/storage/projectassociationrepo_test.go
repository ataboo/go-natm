package storage

import (
	"context"
	"testing"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func _assertAssociationExists(t *testing.T, projectID string, userID string, email string, associationType string) *models.ProjectAssociation {
	user, err := models.FindUser(testCtx, testDb, userID)
	if err != nil {
		t.Errorf("failed to find user with id %s", userID)
		return nil
	}

	association, err := models.ProjectAssociations(qm.Where("project_id = ? AND email = ?", projectID, user.Email)).One(context.Background(), testDb)
	if err != nil {
		t.Errorf("expected project association with user_id %s and project_id %s to exist", userID, projectID)
		return association
	}

	if association.Email != email {
		t.Errorf("unnexpected project association email %s, %s", email, association.Email)
	}

	if association.Association != associationType {
		t.Errorf("unnexpected association type %s, %s", associationType, association.Association)
	}

	return association
}

func _assertAssociationDoesNotExist(t *testing.T, projectID string, userID string) {
	user, err := models.FindUser(testCtx, testDb, userID)
	if err != nil {
		t.Errorf("failed to find user with id %s", userID)
		return
	}

	_, err = models.ProjectAssociations(qm.Where("project_id = ? AND email = ?", projectID, user.Email)).One(context.Background(), testDb)
	if err == nil {
		t.Errorf("expected project association with user_id %s and project_id %s not to exist", userID, projectID)
	}
}

func TestProjectAssociationBasics(t *testing.T) {
	repo_testsetup(t)
	defer repo_testteardown()

	assocRepo := NewProjectAssociationRepository(testDb)
	projRepo := NewProjectRepository(testDb)
	userRepo := NewUserRepository(testDb)

	owner := &models.User{
		Email:  "test@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Project Owner",
	}
	if err := userRepo.Create(owner); err != nil {
		t.Error(err)
	}

	readerUser := &models.User{
		Email:  "reader@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Reader User",
	}
	if err := userRepo.Create(readerUser); err != nil {
		t.Error(err)
	}

	writerUser := &models.User{
		Email:  "writer@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Writer User",
	}
	if err := userRepo.Create(writerUser); err != nil {
		t.Error(err)
	}

	if err := projRepo.Create(&data.ProjectCreate{
		Name:         "Test Project",
		Abbreviation: "PRJ",
		Description:  "A test project",
	}, owner.ID); err != nil {
		t.Error(err)
	}

	testProj, err := models.Projects(qm.Where("abbreviation = ?", "PRJ")).One(context.Background(), testDb)
	if err != nil {
		t.Error(err)
	}

	_ = _assertAssociationExists(t, testProj.ID, owner.ID, owner.Email, models.AssociationsEnumOwner)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, readerUser.ID)

	if err == nil {
		t.Error("readerUser should not be able to create a project association")
	}

	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}

	_assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     writerUser.Email,
		Type:      models.AssociationsEnumWriter,
	}, readerUser.ID)
	if err == nil {
		t.Error("expected error when reader creates association")
	}

	_assertAssociationDoesNotExist(t, testProj.ID, writerUser.ID)

}

// func TestFoo(t *testing.T) {

// 	user := models.User{
// 		Email:  "test@email.com",
// 		Active: true,
// 		ID:     uuid.New().String(),
// 		Name:   "Test User",
// 	}
// 	if err := user.Insert(context.Background(), testDb, boil.Infer()); err != nil {
// 		t.Error(err)
// 	}

// 	project := models.Project{
// 		ID:     uuid.New().String(),
// 		Name:   "Test projects",
// 		Active: true,
// 	}

// 	if err := project.Insert(context.Background(), testDb, boil.Infer()); err != nil {
// 		t.Error(err)
// 	}

// 	err := repo.Create(&data.ProjectAssociationCreate{
// 		ProjectID: project.ID,
// 	}, user.ID)

// 	t.Error(err)
// }

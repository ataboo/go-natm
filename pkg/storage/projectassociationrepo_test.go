package storage

import (
	"context"
	"testing"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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

	otherUser := &models.User{
		Email:  "test2@email.com",
		Active: true,
		ID:     uuid.New().String(),
		Name:   "Other User",
	}
	if err := userRepo.Create(otherUser); err != nil {
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

	ownerAssoc, err := assocRepo.assertActingUserCanChangeAssociations(owner.ID, testProj.ID)
	if err != nil {
		t.Error(err)
	}

	if ownerAssoc.Association != models.AssociationsEnumOwner {
		t.Error("unnexpected association type")
	}

	if ownerAssoc.Email != owner.Email {
		t.Errorf("unnexpected owner email in association")
	}

	otherUserAssoc, err := assocRepo.assertActingUserCanChangeAssociations(otherUser.ID, testProj.ID)
	if err == nil || otherUserAssoc != nil {
		t.Error("expected error")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     otherUser.Email,
		Type:      models.AssociationsEnumReader,
	}, otherUser.ID)

	if err == nil {
		t.Error("expected error")
	}

	if count, _ := models.ProjectAssociations(qm.Where("project_id = ?", testProj.ID)).Count(context.Background(), testDb); count != 1 {
		t.Error("unnexpected assoc count")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     otherUser.Email,
		Type:      models.AssociationsEnumReader,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}

	if count, _ := models.ProjectAssociations(qm.Where("project_id = ?", testProj.ID)).Count(context.Background(), testDb); count != 2 {
		t.Error("unnexpected association count")
	}

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

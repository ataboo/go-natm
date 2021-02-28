package storage

import (
	"runtime"
	"testing"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func _assertAssociationExists(t *testing.T, projectID string, userID string, email string, associationType string) *models.ProjectAssociation {
	_, file, no, _ := runtime.Caller(1)

	user, err := models.FindUser(testCtx, testTx, userID)
	if err != nil {
		t.Errorf("failed to find user with id %s", userID)
		t.Errorf("%s:%d", file, no)
		return nil
	}

	association, err := models.ProjectAssociations(qm.Where("project_id = ? AND email = ?", projectID, user.Email)).One(testCtx, testTx)
	if err != nil {
		t.Errorf("expected project association with user_id %s and project_id %s to exist", userID, projectID)
		t.Errorf("%s:%d", file, no)
		return association
	}

	if association.Email != email {
		t.Errorf("unnexpected project association email %s, %s", email, association.Email)
		t.Errorf("%s:%d", file, no)
	}

	if association.Association != associationType {
		t.Errorf("unnexpected association type %s, %s", associationType, association.Association)
		t.Errorf("%s:%d", file, no)
	}

	return association
}

func _assertAssociationDoesNotExist(t *testing.T, projectID string, userID string) {
	_, file, no, _ := runtime.Caller(1)

	user, err := models.FindUser(testCtx, testTx, userID)
	if err != nil {
		t.Errorf("failed to find user with id %s", userID)
		t.Errorf("%s:%d", file, no)
		return
	}

	_, err = models.ProjectAssociations(qm.Where("project_id = ? AND email = ?", projectID, user.Email)).One(testCtx, testTx)
	if err == nil {
		t.Errorf("expected project association with user_id %s and project_id %s not to exist", userID, projectID)
		t.Errorf("%s:%d", file, no)
	}
}

func TestProjectAssociationBasics(t *testing.T) {
	repo_testsetup(t)
	defer repo_testteardown()

	assocRepo := NewProjectAssociationRepository(testTx)
	projRepo := NewProjectRepository(testTx)
	userRepo := NewUserRepository(testCtx, testTx)

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

	testProj, err := models.Projects(qm.Where("abbreviation = ?", "PRJ")).One(testCtx, testTx)
	if err != nil {
		t.Error(err)
	}

	//===Owner Actions===
	ownerAssoc := _assertAssociationExists(t, testProj.ID, owner.ID, owner.Email, models.AssociationsEnumOwner)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: uuid.New().String(),
		Email:     owner.Email,
		Type:      models.AssociationsEnumOwner,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to add an association for a non-existant project")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     owner.Email,
		Type:      models.AssociationsEnumOwner,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to add a duplicate association")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumOwner,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to create another owner association")
	}
	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     writerUser.Email,
		Type:      models.AssociationsEnumWriter,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	writerAssoc := _assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	readerAssoc := _assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   uuid.New().String(),
		Type: models.AssociationsEnumReader,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to change non-existant associations")
	}

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   writerAssoc.ID,
		Type: "notanenumvalue",
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to set invalid association type")
	}

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   ownerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to change their own association")
	}
	_assertAssociationExists(t, testProj.ID, owner.ID, owner.Email, models.AssociationsEnumOwner)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   readerAssoc.ID,
		Type: models.AssociationsEnumWriter,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   readerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: ownerAssoc.ID,
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to delete their own association")
	}
	_assertAssociationExists(t, testProj.ID, owner.ID, owner.Email, models.AssociationsEnumOwner)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: uuid.New().String(),
	}, owner.ID)
	if err == nil {
		t.Error("owner should not be able to delete non-existant associations")
	}

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: writerAssoc.ID,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationDoesNotExist(t, testProj.ID, writerUser.ID)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: readerAssoc.ID,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	//===Writer Actions===
	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, writerUser.ID)
	if err == nil {
		t.Error("writer without association should not be able to create an association")
	}
	_assertAssociationDoesNotExist(t, testProj.ID, writerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     writerUser.Email,
		Type:      models.AssociationsEnumWriter,
	}, owner.ID)
	if err != nil {
		t.Error(err)
	}
	writerAssoc = _assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumOwner,
	}, writerUser.ID)
	if err == nil {
		t.Error("writer should not be able to create an owner association")
	}
	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumWriter,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	readerAssoc = _assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: readerAssoc.ID,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	readerAssoc = _assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: readerAssoc.ID,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationDoesNotExist(t, testProj.ID, readerUser.ID)

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     readerUser.Email,
		Type:      models.AssociationsEnumReader,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	readerAssoc = _assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: ownerAssoc.ID,
	}, writerUser.ID)
	if err == nil {
		t.Error("writer should not be able to delete owner association")
	}
	_assertAssociationExists(t, testProj.ID, owner.ID, owner.Email, models.AssociationsEnumOwner)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   readerAssoc.ID,
		Type: models.AssociationsEnumWriter,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	readerAssoc = _assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   writerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, writerUser.ID)
	if err == nil {
		t.Error("writer should not be able to change their own association")
	}
	writerAssoc = _assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: writerAssoc.ID,
	}, writerUser.ID)
	if err == nil {
		t.Error("writer should not be able to delete their own association")
	}
	_assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   readerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, writerUser.ID)
	if err != nil {
		t.Error(err)
	}
	_assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	//===ReaderActions===
	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     "someuser@email.com",
		Type:      models.AssociationsEnumReader,
	}, readerUser.ID)
	if err == nil {
		t.Error("reader should not be able to create an association")
	}
	_assertAssociationExists(t, testProj.ID, readerUser.ID, readerUser.Email, models.AssociationsEnumReader)

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   writerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, readerUser.ID)
	if err == nil {
		t.Error("reader should not be able to update an association")
	}
	_assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: writerAssoc.ID,
	}, readerUser.ID)
	if err == nil {
		t.Error("reader should not be able to delete an association")
	}
	_assertAssociationExists(t, testProj.ID, writerUser.ID, writerUser.Email, models.AssociationsEnumWriter)

	//===Invalid User Actions===
	invalidUserID := uuid.New().String()
	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     "invalid@email.com",
		Type:      models.AssociationsEnumOwner,
	}, invalidUserID)
	if err == nil {
		t.Error("invalid user id should not be able to create an owner association")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     "invalid@email.com",
		Type:      models.AssociationsEnumWriter,
	}, invalidUserID)
	if err == nil {
		t.Error("invalid user id should not be able to create an writer association")
	}

	err = assocRepo.Create(&data.ProjectAssociationCreate{
		ProjectID: testProj.ID,
		Email:     "invalid@email.com",
		Type:      models.AssociationsEnumReader,
	}, invalidUserID)
	if err == nil {
		t.Error("invalid user id should not be able to create an reader association")
	}

	err = assocRepo.Delete(&data.ProjectAssociationDelete{
		ID: writerAssoc.ID,
	}, invalidUserID)
	if err == nil {
		t.Error("invalid user should not be able to delete an association")
	}

	err = assocRepo.Update(&data.ProjectAssociationUpdate{
		ID:   writerAssoc.ID,
		Type: models.AssociationsEnumReader,
	}, invalidUserID)
	if err == nil {
		t.Error("invalid user should not be able to update an association")
	}
}

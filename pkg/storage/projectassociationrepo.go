package storage

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProjectAssociationRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewProjectAssociationRepository(db *sql.DB) *ProjectAssociationRepository {
	return &ProjectAssociationRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *ProjectAssociationRepository) Create(createData *data.ProjectAssociationCreate, actingUserID string) error {
	_, err := r.assertActingUserCanChangeAssociations(actingUserID, createData.ProjectID)
	if err != nil {
		return err
	}

	project, err := models.Projects(qm.Where("id = ?", createData.ProjectID), qm.Load("ProjectAssociations")).One(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	for _, association := range project.R.ProjectAssociations {
		if association.Email == createData.Email {
			return &common.ErrorWithStatus{Code: http.StatusBadRequest}
		}
	}

	if createData.Type != models.AssociationsEnumReader && createData.Type != models.AssociationsEnumWriter {
		return &common.ErrorWithStatus{Code: http.StatusBadRequest}
	}

	projAssociation := models.ProjectAssociation{
		ID:          uuid.New().String(),
		Association: createData.Type,
		ProjectID:   createData.ProjectID,
		Email:       createData.Email,
	}

	err = projAssociation.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *ProjectAssociationRepository) Update(updateData *data.ProjectAssociationUpdate, actingUserID string) error {
	subjectAssociation, err := models.FindProjectAssociation(r.ctx, r.db, updateData.ID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	if subjectAssociation.Association == models.AssociationsEnumOwner {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	actingUserAssociation, err := r.assertActingUserCanChangeAssociations(actingUserID, subjectAssociation.ProjectID)
	if err != nil {
		return err
	}

	if actingUserAssociation.Email == subjectAssociation.Email {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	if updateData.Type != models.AssociationsEnumReader && updateData.Type != models.AssociationsEnumWriter {
		return &common.ErrorWithStatus{Code: http.StatusBadRequest}
	}

	subjectAssociation.Association = updateData.Type
	_, err = subjectAssociation.Update(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *ProjectAssociationRepository) Delete(deleteData *data.ProjectAssociationDelete, actingUserID string) error {
	subjectAssociation, err := models.FindProjectAssociation(r.ctx, r.db, deleteData.ID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	if subjectAssociation.Association == models.AssociationsEnumOwner {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	actingUserAssociation, err := r.assertActingUserCanChangeAssociations(actingUserID, subjectAssociation.ProjectID)
	if err != nil {
		return err
	}

	if actingUserAssociation.Email == subjectAssociation.Email {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	_, err = subjectAssociation.Delete(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *ProjectAssociationRepository) assertActingUserCanChangeAssociations(actingUserID string, projectID string) (*models.ProjectAssociation, error) {
	actingUser, err := models.FindUser(r.ctx, r.db, actingUserID)
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	actingUserAssociation, err := models.ProjectAssociations(
		qm.Where("email = ? AND project_id = ?", actingUser.Email, projectID),
	).One(r.ctx, r.db)
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	// Only owners and writers can set associations
	if actingUserAssociation.Association != models.AssociationsEnumOwner && actingUserAssociation.Association != models.AssociationsEnumWriter {
		return actingUserAssociation, &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	return actingUserAssociation, nil
}

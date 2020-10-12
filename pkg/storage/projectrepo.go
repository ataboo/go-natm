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

type ProjectRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *ProjectRepository) ListAssociated(userID string) (models.ProjectAssociationSlice, error) {
	qms := []qm.QueryMod{
		qm.Where("user_id = ?", userID),
		qm.Load("Project"),
	}

	return models.ProjectAssociations(qms...).All(r.ctx, r.db)
}

func (r *ProjectRepository) Find(projectID string, userID string) (*models.ProjectAssociation, error) {
	return models.ProjectAssociations(
		qm.Where("user_id = ? AND project_id = ?", userID, projectID),
		qm.Load("Project"),
		qm.Load("Project.TaskStatuses.Tasks"),
		qm.Load("Project.TaskStatuses.Tasks.WorkLogs"),
		qm.Load("Project.TaskStatuses.Tasks.Assignee"),
		qm.Load("Project.ProjectAssociations.User"),
	).One(r.ctx, r.db)
}

func (r *ProjectRepository) Create(projectData *data.ProjectCreate, ownerID string) error {
	proj := models.Project{
		ID:           uuid.New().String(),
		Abbreviation: projectData.Abbreviation,
		Description:  projectData.Description,
		Active:       true,
		Name:         projectData.Name,
	}

	projAssociation := models.ProjectAssociation{
		ID:          uuid.New().String(),
		Association: models.AssociationsEnumOwner,
		ProjectID:   proj.ID,
		UserID:      ownerID,
	}

	err := proj.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return projAssociation.Insert(r.ctx, r.db, boil.Infer())
}

func (r *ProjectRepository) Update(projectData *data.ProjectUpdate) error {
	proj := models.Project{
		ID:           projectData.ID,
		Active:       projectData.Active,
		Abbreviation: projectData.Abbreviation,
		Description:  projectData.Description,
		Name:         projectData.Name,
	}

	_, err := proj.Update(r.ctx, r.db, boil.Infer())

	return err
}

func (r *ProjectRepository) Archive(projectID string, userID string) error {
	association, err := r.Find(projectID, userID)
	if err != nil || association == nil {
		return &common.ErrorWithStatus{
			Code: http.StatusNotFound,
		}
	}

	switch association.Association {
	case models.AssociationsEnumOwner:
	case models.AssociationsEnumWriter:
		break
	default:
		return &common.ErrorWithStatus{
			Code: http.StatusForbidden,
		}
	}

	association.R.Project.Active = false

	_, err = association.R.Project.Update(r.ctx, r.db, boil.Infer())

	if err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}

func (r *ProjectRepository) SetTaskOrder(taskOrder data.ProjectTaskOrder, userID string) error {
	association, err := r.Find(taskOrder.ID, userID)
	if err != nil || association == nil {
		return &common.ErrorWithStatus{
			Code: http.StatusNotFound,
		}
	}

	switch association.Association {
	case models.AssociationsEnumOwner:
	case models.AssociationsEnumWriter:
		break
	default:
		return &common.ErrorWithStatus{
			Code: http.StatusForbidden,
		}
	}

	taskMap := make(map[string]data.TaskOrder)
	for _, t := range taskOrder.Tasks {
		taskMap[t.ID] = t
	}

	tasks, err := models.Tasks(qm.LeftOuterJoin("task_statuses s ON s.id = tasks.task_status_id"), qm.Where("s.project_id = ?", taskOrder.ID)).All(r.ctx, r.db)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		newOrder, ok := taskMap[t.ID]
		if !ok {
			return err
		}

		if t.TaskStatusID != newOrder.StatusID || t.Ordinal != newOrder.Ordinal {
			t.TaskStatusID = newOrder.StatusID
			t.Ordinal = newOrder.Ordinal
			_, err := t.Update(r.ctx, r.db, boil.Infer())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

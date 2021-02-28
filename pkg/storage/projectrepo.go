package storage

import (
	"context"
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//ProjectRepository handles changes to Project models.
type ProjectRepository struct {
	db  boil.ContextExecutor
	ctx context.Context
}

//NewProjectRepository creates a new project repository
func NewProjectRepository(db boil.ContextExecutor) *ProjectRepository {
	return &ProjectRepository{
		db:  db,
		ctx: context.Background(),
	}
}

//ListAssociated lists all project associations for the projects that the user is associated with.
func (r *ProjectRepository) ListAssociated(userID string) (models.ProjectAssociationSlice, error) {
	actingUser, err := models.FindUser(r.ctx, r.db, userID)
	if err != nil {
		return models.ProjectAssociationSlice{}, &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	qms := []qm.QueryMod{
		qm.Where("email = ?", actingUser.Email),
		qm.Load("Project", qm.Select("name")),
	}

	return models.ProjectAssociations(qms...).All(r.ctx, r.db)
}

//Find finds a project by id.
func (r *ProjectRepository) Find(projectID string, userID string) (*models.ProjectAssociation, error) {
	actingUser, err := models.FindUser(r.ctx, r.db, userID)
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	return models.ProjectAssociations(
		qm.Where("email = ? AND project_id = ?", actingUser.Email, projectID),
		qm.Load("Project"),
		qm.Load("Project.TaskStatuses.Tasks"),
		qm.Load("Project.TaskStatuses.Tasks.WorkLogs"),
		qm.Load("Project.TaskStatuses.Tasks.Assignee"),
		qm.Load("Project.ProjectAssociations", qm.OrderBy("email")),
	).One(r.ctx, r.db)
}

//Create creates a new project owned by the acting user.
func (r *ProjectRepository) Create(projectData *data.ProjectCreate, actingUserID string) error {
	actingUser, err := models.FindUser(r.ctx, r.db, actingUserID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	if projectData.Abbreviation == "" || projectData.Description == "" || projectData.Name == "" {
		return &common.ErrorWithStatus{Code: http.StatusUnprocessableEntity}
	}

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
		Email:       actingUser.Email,
	}

	err = proj.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return projAssociation.Insert(r.ctx, r.db, boil.Infer())
}

//Update updates a projects.
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

//Archive soft-deletes a project.
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

//SetTaskOrder sets the order of tasks within a project.
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

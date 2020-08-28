package storage

import (
	"context"
	"database/sql"

	"github.com/ataboo/go-natm/v4/pkg/api/data"
	"github.com/ataboo/go-natm/v4/pkg/models"
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
	return models.ProjectAssociations(
		qm.Where("user_id = ?", userID),
		qm.Load("Project"),
	).All(r.ctx, r.db)
}

func (r *ProjectRepository) Find(projectID string, userID string) (*models.ProjectAssociation, error) {
	return models.ProjectAssociations(
		qm.Where("user_id = ? AND project_id = ?", userID, projectID),
		qm.Load("Project"),
		qm.Load("Project.Tasks"),
		qm.Load("Project.Tasks.WorkLogs"),
	).One(r.ctx, r.db)
}

func (r *ProjectRepository) Create(projectData *data.ProjectCreate) error {
	proj := models.Project{
		Identifier: projectData.Identifier,
		Active:     true,
		Name:       projectData.Name,
	}

	return proj.Insert(r.ctx, r.db, boil.Infer())
}

func (r *ProjectRepository) Update(projectData *data.ProjectUpdate) error {
	proj := models.Project{
		ID:         projectData.ID,
		Active:     projectData.Active,
		Identifier: projectData.Identifier,
		Name:       projectData.Name,
	}

	_, err := proj.Update(r.ctx, r.db, boil.Infer())

	return err
}

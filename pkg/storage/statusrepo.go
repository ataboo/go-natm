package storage

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type StatusRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewStatusRepository(db *sql.DB) *StatusRepository {
	return &StatusRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *StatusRepository) Archive(statusID string, userID string) error {
	actingUser, err := models.FindUser(r.ctx, r.db, userID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	status, err := models.TaskStatuses(
		qm.Load("Tasks"),
		qm.LeftOuterJoin("project_associations pa ON pa.project_id = task_statuses.project_id"),
		qm.Where("pa.email = ? AND task_statuses.id = ?", actingUser.Email, statusID),
	).One(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	otherStatuses, err := models.TaskStatuses(
		qm.Where("id != ? AND project_id = ? AND active = TRUE", statusID, status.ProjectID),
		qm.OrderBy("ordinal"),
	).All(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	var taskUpdateCols models.M
	if len(otherStatuses) > 0 {
		taskUpdateCols = models.M{"task_status_id": otherStatuses[0].ID}
	} else {
		taskUpdateCols = models.M{"active": false}
	}

	_, err = status.R.Tasks.UpdateAll(r.ctx, r.db, taskUpdateCols)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	status.Active = false
	_, err = status.Update(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	if err = r.normalizeStatusOrdinals(status.ProjectID); err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *StatusRepository) StepOrdinal(statusID string, userID string, step int) error {
	actingUser, err := models.FindUser(r.ctx, r.db, userID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	status, err := models.TaskStatuses(
		qm.LeftOuterJoin("project_associations pa ON pa.project_id = task_statuses.project_id"),
		qm.Where("pa.email = ? AND task_statuses.id = ?", actingUser.Email, statusID),
	).One(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusNotFound,
		}
	}

	statuses, err := models.TaskStatuses(
		qm.Where("task_statuses.active = TRUE AND task_statuses.project_id = ?", status.ProjectID),
		qm.OrderBy("ordinal"),
	).All(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusInternalServerError,
		}
	}

	swappedStatusOrdinal := status.Ordinal + step
	if swappedStatusOrdinal < 0 || swappedStatusOrdinal >= len(statuses) {
		return &common.ErrorWithStatus{
			Code: http.StatusBadRequest,
		}
	}

	swappedStatus := statuses[swappedStatusOrdinal]
	swappedStatus.Ordinal = status.Ordinal
	status.Ordinal = swappedStatusOrdinal
	if _, err = swappedStatus.Update(r.ctx, r.db, boil.Infer()); err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusInternalServerError,
		}
	}

	if _, err = status.Update(r.ctx, r.db, boil.Infer()); err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}

func (r *StatusRepository) Create(data *data.StatusCreate, userID string) error {
	maxOrdinal := struct {
		MaxOrdinal sql.NullInt32 `boil:"max_ordinal" json:"max_ordinal" toml:"max_ordinal" yaml:"max_ordinal"`
	}{}

	query := queries.Raw(
		`SELECT MAX(s.ordinal) max_ordinal
		FROM projects
		LEFT JOIN task_statuses s ON projects.id = s.project_id
		WHERE projects.id = $1`,
		data.ProjectID,
	)

	err := query.Bind(r.ctx, r.db, &maxOrdinal)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	nextOrdinal := 0
	if maxOrdinal.MaxOrdinal.Valid {
		nextOrdinal = int(maxOrdinal.MaxOrdinal.Int32) + 1
	}

	statusModel := models.TaskStatus{
		ID:        uuid.New().String(),
		Active:    true,
		Name:      data.Name,
		Ordinal:   nextOrdinal,
		ProjectID: data.ProjectID,
	}

	return statusModel.Insert(r.ctx, r.db, boil.Infer())
}

func (r *StatusRepository) normalizeStatusOrdinals(projectID string) error {
	statuses, err := models.TaskStatuses(
		qm.Where("project_id = ? AND active = true", projectID),
		qm.OrderBy("ordinal"),
	).All(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{
			Code: http.StatusNotFound,
		}
	}

	for i, s := range statuses {
		s.Ordinal = i
		if _, err := s.Update(r.ctx, r.db, boil.Infer()); err != nil {
			return &common.ErrorWithStatus{
				Code: http.StatusInternalServerError,
			}
		}
	}

	return nil
}

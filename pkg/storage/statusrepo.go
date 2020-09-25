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
	status, err := models.TaskStatuses(
		qm.Load("Tasks"),
		qm.LeftOuterJoin("project_associations pa ON pa.project_id = task_statuses.project_id"),
		qm.Where("pa.user_id = ? AND task_statuses.id = ?", userID, statusID),
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

	for _, t := range status.R.Tasks {
		if (len(otherStatuses) > 0) {
			otherStatuses[0].ID
		} else {
			
		}	
	}

	var nextStatus &models.TaskStatus = nil
	for _, status := range otherStatuses {
		if nextStatus == nil || status.Ordinal > nextStatus.
	}

	status.Active = false
	_, err = status.Update(r.ctx, r.db, boil.Infer())

	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
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

	// // .Bind(r.ctx, r.db, &maxOrdinal)
	// if err != nil {
	// 	return err
	// }

	statusModel := models.TaskStatus{
		ID:        uuid.New().String(),
		Active:    true,
		Name:      data.Name,
		Ordinal:   int(maxOrdinal.MaxOrdinal.Int32) + 1,
		ProjectID: data.ProjectID,
	}

	return statusModel.Insert(r.ctx, r.db, boil.Infer())
}

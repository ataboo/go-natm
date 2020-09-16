package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
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

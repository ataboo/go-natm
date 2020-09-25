package storage

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskRepository struct {
	db       *sql.DB
	ctx      context.Context
	dayExpr  *regexp.Regexp
	hourExpr *regexp.Regexp
}

func NewTaskRepository(db *sql.DB) *TaskRepository {

	dayExpr, _ := regexp.Compile(`(\d*\.?\d*)[dD]`)
	hourExpr, _ := regexp.Compile(`(\d*\.?\d*)[hH]`)

	return &TaskRepository{
		db:       db,
		ctx:      context.Background(),
		dayExpr:  dayExpr,
		hourExpr: hourExpr,
	}
}

func (r *TaskRepository) List(projectID string) (models.TaskSlice, error) {
	return models.Tasks(
		qm.Where("project_id = ?", projectID),
		qm.Where("active = ?", true),
	).All(r.ctx, r.db)
}

func (r *TaskRepository) Find(taskID string, userID string) (*models.Task, error) {
	return models.Tasks(
		qm.LeftOuterJoin("task_statuses s ON s.id = tasks.task_status_id"),
		qm.LeftOuterJoin("project_associations pa ON pa.project_id = s.project_id"),
		qm.Where("tasks.id = ? AND pa.user_id = ?", taskID, userID),
	).One(r.ctx, r.db)
}

func (r *TaskRepository) Archive(taskID string, userID string) error {
	taskModel, err := r.Find(taskID, userID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	taskModel.Active = false

	_, err = taskModel.Update(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *TaskRepository) Create(taskData *data.TaskCreate, ownerID string) error {
	statusModel, err := models.FindTaskStatus(r.ctx, r.db, taskData.StatusID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	maxValRow := struct {
		MaxValue sql.NullInt32 `boil:"max_value" json:"max_value" toml:"max_value" yaml:"max_value"`
	}{}

	err = queries.Raw(
		`SELECT MAX(t.number) max_value
		from task_statuses
		LEFT JOIN tasks t ON task_statuses.id = t.task_status_id
		WHERE task_statuses.project_id = $1`,
		statusModel.ProjectID,
	).Bind(r.ctx, r.db, &maxValRow)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	maxNumber := int(maxValRow.MaxValue.Int32)

	err = queries.Raw(
		`SELECT MAX(tasks.ordinal) max_value
		FROM tasks
		WHERE tasks.task_status_id = $1`,
		taskData.StatusID,
	).Bind(r.ctx, r.db, &maxValRow)
	if err != nil {
		return err
	}
	maxOrdinal := int(maxValRow.MaxValue.Int32)

	taskModel := models.Task{
		AssigneeID:   null.NewString(taskData.AssigneeID, taskData.AssigneeID != ""),
		Description:  taskData.Description,
		ID:           uuid.New().String(),
		Number:       maxNumber + 1,
		Ordinal:      maxOrdinal + 1,
		TaskStatusID: taskData.StatusID,
		TaskType:     taskData.Type,
		Title:        taskData.Title,
	}

	return taskModel.Insert(r.ctx, r.db, boil.Infer())
}

func (r *TaskRepository) Update(taskData *data.TaskUpdate, userID string) error {
	var assigneeID = ""
	user, err := models.Users(qm.Where("email = ?", taskData.AssigneeEmail)).One(r.ctx, r.db)
	if err == nil && user != nil {
		assigneeID = user.ID
	}

	estimateMins, ok := r.parseDuration(taskData.Estimate)

	task, err := models.FindTask(r.ctx, r.db, taskData.ID)
	task.AssigneeID = null.NewString(assigneeID, assigneeID != "")
	task.Description = taskData.Description
	task.Estimate = null.NewInt(estimateMins, ok)
	task.ID = taskData.ID
	task.Title = taskData.Title
	task.TaskType = taskData.Type

	_, err = task.Update(r.ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) parseDuration(input string) (duration int, ok bool) {
	if match := r.dayExpr.FindStringSubmatch(input); len(match) == 2 {
		if flDays, err := strconv.ParseFloat(match[1], 32); err == nil {
			return int(math.Round(flDays * 8 * 60)), true
		}
	}

	if match := r.hourExpr.FindStringSubmatch(input); len(match) == 2 {
		if flHours, err := strconv.ParseFloat(match[1], 32); err == nil {
			return int(math.Round(flHours * 60)), true
		}
	}

	input = strings.Replace(input, "m", "", 1)
	if flMins, err := strconv.ParseFloat(input, 32); err == nil {
		return int(math.Round(flMins)), true
	}

	return 0, false
}

// func (r *TaskRepository) Archive(projectID string, userID string) error {
// 	association, err := r.Find(projectID, userID)
// 	if err != nil || association == nil {
// 		return &common.ErrorWithStatus{
// 			Code: http.StatusNotFound,
// 		}
// 	}

// 	switch association.Association {
// 	case models.AssociationsEnumOwner:
// 	case models.AssociationsEnumWriter:
// 		break
// 	default:
// 		return &common.ErrorWithStatus{
// 			Code: http.StatusForbidden,
// 		}
// 	}

// 	association.R.Project.Active = false

// 	_, err = association.R.Project.Update(r.ctx, r.db, boil.Infer())

// 	if err != nil {
// 		return &common.ErrorWithStatus{
// 			Code: http.StatusInternalServerError,
// 		}
// 	}

// 	return nil
// }

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
	"time"

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
	task.Estimate = null.NewInt(estimateMins*60, ok)
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

func (r *TaskRepository) StartLoggingWork(userID string, taskID string) error {
	task, err := r.Find(taskID, userID)
	if err != nil || task == nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	err = r.StopLoggingWork(userID)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	newLog := models.WorkLog{
		ID:        uuid.New().String(),
		StartTime: time.Now().UTC(),
		TaskID:    taskID,
		UserID:    userID,
	}

	err = newLog.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *TaskRepository) StopLoggingWork(userID string) error {
	openWorkLogs, err := models.WorkLogs(
		qm.Where("work_logs.user_id = ? AND work_logs.end_time IS NULL", userID),
	).All(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	_, err = openWorkLogs.UpdateAll(r.ctx, r.db, models.M{"end_time": time.Now().UTC()})
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

func (r *TaskRepository) AddComment(userID string, commentCreate *data.CommentCreate) (*data.CommentRead, error) {
	_, err := r.Find(commentCreate.TaskID, userID)
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	user, err := models.FindUser(r.ctx, r.db, userID)
	if user == nil || err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	comment := models.Comment{
		ID:      uuid.New().String(),
		TaskID:  commentCreate.TaskID,
		UserID:  userID,
		Message: commentCreate.Message,
	}

	err = comment.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	commentVM := &data.CommentRead{
		Author: &data.UserRead{
			Email: user.Email,
			ID:    user.ID,
			Name:  user.Name,
		},
		ID:        comment.ID,
		Message:   commentCreate.Message,
		TaskID:    commentCreate.TaskID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return commentVM, nil
}

func (r *TaskRepository) GetComments(userID string, taskID string) ([]data.CommentRead, error) {
	task, err := r.Find(taskID, userID)
	if err != nil || task == nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	comments, err := models.Comments(qm.Where("task_id = ?", taskID), qm.Load("User"), qm.OrderBy("created_at")).All(r.ctx, r.db)
	if err != nil {
		return nil, &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	commentVMs := make([]data.CommentRead, len(comments))
	for i, c := range comments {
		commentVMs[i] = *&data.CommentRead{
			Author: &data.UserRead{
				Email: c.R.User.Email,
				ID:    c.R.User.ID,
				Name:  c.R.User.Name,
			},
			CreatedAt: c.CreatedAt.UTC(),
			UpdatedAt: c.UpdatedAt.UTC(),
			Message:   c.Message,
			TaskID:    c.TaskID,
			ID:        c.ID,
		}
	}

	return commentVMs, nil
}

func (r *TaskRepository) DeleteComment(userID string, commentID string) error {
	comment, err := models.Comments(qm.Where("id = ?", commentID)).One(r.ctx, r.db)
	if err != nil {
		return &common.ErrorWithStatus{Code: http.StatusNotFound}
	}

	if comment.UserID != userID {
		return &common.ErrorWithStatus{Code: http.StatusForbidden}
	}

	if _, err = comment.Delete(r.ctx, r.db); err != nil {
		return &common.ErrorWithStatus{Code: http.StatusInternalServerError}
	}

	return nil
}

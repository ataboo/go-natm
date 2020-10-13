package data

import (
	"github.com/volatiletech/null/v8"
)

type ProjectGrid struct {
	ID              string `json:"id"`
	AssociationType string `json:"associationType"`
	Name            string `json:"name"`
	Abbreviation    string `json:"abbreviation"`
	Description     string `json:"description"`
	LastUpdated     int64  `json:"lastUpdated"`
}

type ProjectDetail struct {
	ID            string                     `json:"id"`
	Name          string                     `json:"name"`
	Abbreviation  string                     `json:"abbreviation"`
	Description   string                     `json:"description"`
	Statuses      []StatusRead               `json:"statuses"`
	Tasks         []TaskGrid                 `json:"tasks"`
	WorkingTaskID null.String                `json:"workingTaskID"`
	Associations  []ProjectAssociationDetail `json:"associations"`
}

type ProjectCreate struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
	Description  string `json:"description" binding:"required"`
}

type ProjectUpdate struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"identifier" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Active       bool   `json:"active" binding:"required"`
}

type ProjectArchive struct {
	ProjectID string `json:"projectID" binding:"required"`
}

type ProjectTaskOrder struct {
	ID    string      `json:"id" binding:"required"`
	Tasks []TaskOrder `json:"tasks" binding:"required"`
}

type TaskOrder struct {
	ID       string `json:"id" binding:"required"`
	StatusID string `json:"statusId" binding:"required"`
	Ordinal  int    `json:"ordinal" binding:"required"`
}

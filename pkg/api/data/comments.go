package data

import "time"

type CommentCreate struct {
	Message string `json:"message"`
	TaskID  string `json:"taskId"`
}

type CommentRead struct {
	Message   string    `json:"message"`
	TaskID    string    `json:"taskId"`
	Author    *UserRead `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

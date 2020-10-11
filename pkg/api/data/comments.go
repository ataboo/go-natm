package data

import "time"

type CommentCreate struct {
	Message string `json:"message"`
	TaskID  string `json:"taskId"`
}

type CommentRead struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	TaskID    string    `json:"taskId"`
	Author    *UserRead `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommentUpdate struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	TaskID  string `json:"taskId"`
}

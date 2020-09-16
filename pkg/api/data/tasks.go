package data

type TimingGrid struct {
	Estimate string `json:"estimate"`
	Current  string `json:"current"`
}

type TaskGrid struct {
	ID          string     `json:"id"`
	Identifier  string     `json:"identifier"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StatusID    string     `json:"statusId"`
	Type        string     `json:"type"`
	Timing      TimingGrid `json:"timing"`
}

type TaskCreate struct {
	AssigneeID  string `json:"assigneeId"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	StatusID    string `json:"statusId" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type TaskUpdate struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	StatusID    string `json:"statusId" binding:"required"`
	Ordinal     int    `json:"ordinal" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

package data

type StatusRead struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StatusCreate struct {
	ProjectID string `json:"projectID"`
	Name      string `json:"name"`
}

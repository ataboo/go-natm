package data

type StatusRead struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Ordinal int    `json:"ordinal"`
}

type StatusCreate struct {
	ProjectID string `json:"projectID"`
	Name      string `json:"name"`
	Ordinal   int    `json:"ordinal"`
}

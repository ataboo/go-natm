package data

type ProjectGrid struct {
	ID              string `json:"id"`
	AssociationType string `json:"association_type"`
	Name            string `json:"name"`
	Abbreviation    string `json:"abbreviation"`
	Description     string `json:"description"`
}

type ProjectDetail struct {
	ID              string       `json:"id"`
	AssociationType string       `json:"association_type"`
	Name            string       `json:"name"`
	Abbreviation    string       `json:"abbreviation"`
	Description     string       `json:"description"`
	Tasks           []TaskGrid   `json:"tasks"`
	Statuses        []StatusRead `json:"statuses"`
}

type ProjectCreate struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
}

type ProjectUpdate struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"identifier"`
	Description  string `json:"description"`
	Active       bool   `json:"active"`
}

package data

type ProjectGrid struct {
	ID              string `json:"id"`
	AssociationType string `json:"association_type"`
	Name            string `json:"name"`
	Identifier      string `json:"identifier"`
}

type ProjectDetail struct {
	ID              string     `json:"id"`
	AssociationType string     `json:"association_type"`
	Name            string     `json:"name"`
	Identifier      string     `json:"identifier"`
	Tasks           []TaskGrid `json:"tasks"`
}

type ProjectCreate struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type ProjectUpdate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Active     bool   `json:"active"`
}

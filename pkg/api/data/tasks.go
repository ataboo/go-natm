package data

type TimingGrid struct {
	Estimate string `json:"estimate"`
	Current  string `json:"current"`
}

type TaskGrid struct {
	ID         string     `json:"id"`
	Identifier string     `json:"identifier"`
	Name       string     `json:"name"`
	StatusID   string     `json:"statusId"`
	Type       string     `json:"type"`
	Timing     TimingGrid `json:"timing"`
}

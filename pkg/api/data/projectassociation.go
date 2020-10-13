package data

type ProjectAssociationDetail struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type ProjectAssociationCreate struct {
	ProjectID string `json:"projectId" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Type      string `json:"type" binding:"required"`
}

type ProjectAssociationUpdate struct {
	ID   string `json:"id" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type ProjectAssociationDelete struct {
	ID string `json:"id" binding:"required"`
}

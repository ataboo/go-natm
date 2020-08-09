package storage

import (
	"database/sql"

	"github.com/ataboo/go-natm/v4/pkg/models"
	"github.com/google/uuid"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	repo := UserRepository{
		db: db,
	}

	return &repo
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Find(id uuid.UUID) *models.User {
	// todo

	return &models.User{}
}

func (r *UserRepository) All() []*models.User {
	//todo

	return nil
}

func (r *UserRepository) FindByEmail(email string) *models.User {
	return &models.User{}
}

func (r *UserRepository) CreateOrUpdate(user *models.User) {
	if user.ID == "" {
		user.ID = uuid.New().String()
		// insert
	} else {
		// update
	}
}

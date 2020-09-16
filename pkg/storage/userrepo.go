package storage

import (
	"context"
	"database/sql"

	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	repo := UserRepository{
		db:  db,
		ctx: context.Background(),
	}

	return &repo
}

type UserRepository struct {
	db  *sql.DB
	ctx context.Context
}

func (r *UserRepository) Find(id string) (*models.User, error) {
	return models.FindUser(r.ctx, r.db, id)
}

func (r *UserRepository) FindByEmail(email string) *models.User {
	user, _ := models.Users(qm.Where("email = ?", email)).One(r.ctx, r.db)

	return user
}

func (r *UserRepository) Create(user *models.User) error {
	user.ID = uuid.New().String()
	err := user.Insert(r.ctx, r.db, boil.Infer())

	return err
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := user.Update(r.ctx, r.db, boil.Infer())

	return err
}

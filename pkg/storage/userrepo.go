package storage

import (
	"context"

	"github.com/ataboo/go-natm/pkg/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//UserRepository handles actions involving User models.
type UserRepository struct {
	db  boil.ContextExecutor
	ctx context.Context
}

//NewUserRepository creates a new UserRepository.
func NewUserRepository(ctx context.Context, db boil.ContextExecutor) *UserRepository {
	repo := UserRepository{
		db:  db,
		ctx: ctx,
	}

	return &repo
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

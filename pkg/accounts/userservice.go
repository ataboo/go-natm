package accounts

import (
	"github.com/ataboo/go-natm/v4/pkg/models"
	"github.com/google/uuid"
)

var enforceInterface UserService = &memUserService{}

type UserService interface {
	GetOrCreateUser(email string, name string) *models.User
}

type memUserService struct {
	users map[uuid.UUID]*models.User
}

func (s *memUserService) GetOrCreateUser(email string, name string) *models.User {
	for _, user := range s.users {
		if user.Name == name {
			return user
		}
	}

	id := uuid.New()
	s.users[id] = &models.User{
		ID:    id.String(),
		Email: email,
		Name:  name,
	}

	return s.users[id]
}

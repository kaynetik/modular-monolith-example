package users

import (
	"github.com/gofrs/uuid"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/config"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/models"
)

type service struct {
	repo Repository
}

type Repository interface {
	GetConfig() *config.Config

	CreateUser(newUser *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
}

func newService(repo Repository) service {
	return service{
		repo,
	}
}

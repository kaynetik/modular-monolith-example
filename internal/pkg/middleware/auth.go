package middleware

import (
	"github.com/kaynetik/modular-monolith-example/internal/pkg/config"

	"github.com/gofrs/uuid"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/models"
)

type userFetcher interface {
	GetConfig() *config.Config

	GetUserByID(uuid.UUID) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
}

func IsAuthenticated() {
	// Placeholder for authentication middleware.
}

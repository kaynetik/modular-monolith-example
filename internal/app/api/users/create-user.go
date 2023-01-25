package users

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kaynetik/modular-monolith-example/internal/app/errors"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/models"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *service) handleCreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newUser models.User

		if err := json.Unmarshal(c.Body(), &newUser); err != nil {
			c.Status(http.StatusUnprocessableEntity)

			return c.JSON(errors.ErrUnprocessableEntity.Map())
		}

		err := s.repo.CreateUser(&newUser)
		if err != nil {
			log.Err(err).Msg("failed to create user")

			return c.SendStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusCreated)

		return c.JSON(newUser)
	}
}

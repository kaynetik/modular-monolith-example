package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (s *service) handleGetUserByEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := s.repo.GetUserByEmail(c.Query("email"))
		if err != nil {
			log.Error().Err(err).Msg("")

			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(user)
	}
}

package users

import (
	"github.com/gofiber/fiber/v2"
)

func Attach(g fiber.Router, repo Repository) {
	svc := newService(repo)

	g.Post("/", svc.handleCreateUser())
	g.Get("/", svc.handleGetUserByEmail())
}

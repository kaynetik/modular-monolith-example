package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kaynetik/modular-monolith-example/internal/app/api/users"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/server"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/storage"
)

// Attach initializes routing information.
func Attach(eng server.Engine, repo *storage.Repository) *server.Engine {
	apiGroup := eng.Get().Group(repo.GetConfig().APIVer)

	apiGroup.Use(corsChain())
	apiGroup.Use(headersChain())

	// Protected Routes for non-admin users.
	//apiGroup.Use(middleware.IsAuthenticated()
	users.Attach(apiGroup.Group("/users"), repo)

	return &eng
}

func corsChain() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH",
		AllowCredentials: true,
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, " +
			"X-CSRF-Token,  X-Requested-With, Access-Control-Allow-Origin" +
			"Access-Control-Allow-Headers, Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
	})
}

func headersChain() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if string(c.Request().Header.Method()) != http.MethodOptions {
			return c.Next()
		}

		// Set security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		return c.Next()
	}
}

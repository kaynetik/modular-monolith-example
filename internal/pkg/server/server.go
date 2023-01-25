package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/config"
	"github.com/rs/zerolog/log"
)

type Engine struct {
	server *fiber.App
	config *config.Server
}

func (e *Engine) New(conf *config.Server) {
	if conf == nil {
		log.Fatal().Msg("config passed to engine initiation was nil")

		return
	}

	e.config = conf

	app := fiber.New(fiber.Config{
		ReadTimeout:  e.config.ReadTimeout,
		WriteTimeout: e.config.WriteTimeout,
	})

	app.Use(logger())

	e.server = app
}

func (e *Engine) Serve() error {
	serverAddr := fmt.Sprintf(":%s", e.config.HTTPPort)

	log.Info().Msgf("server is running on the '%s' port", serverAddr)

	return e.server.Listen(serverAddr)
}

func (e *Engine) Get() *fiber.App {
	return e.server
}

func logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		const (
			statusCode = "status"
			httpMethod = "method"
			path       = "path"
			ip         = "ip"
			reqHeaders = "request_headers"
			reqQuery   = "request_query"
			latency    = "latency"
			userAgent  = "user_agent"
			httpReq    = "http_request"
		)

		start := time.Now()

		// Do not delay the request, let it pass.
		// Deal with rest of the data in an immutable fashion!
		err := c.Next()
		if err != nil {
			log.Error().Err(err).Msg("")

			return err
		}

		code := c.Response().StatusCode()

		// TODO: Expand IP logging to include parsing of the X-Forwarded-For and True-Client-IP headers, as a bear minimum.
		// TODO: Add a config option to enable/disable logging of the request headers and query.
		// TODO: Decrease logging of the request headers to only what would be useful in production.
		reqLogger := log.Logger.With().
			Int(statusCode, code).
			Str(httpMethod, c.Method()).
			Str(path, c.Path()).
			Str(ip, c.IP()). // Note that this will rarely produce the correct IP address.
			Str(reqHeaders, c.Request().Header.String()).
			Str(reqQuery, string(c.Request().URI().QueryString())).
			Str(latency, time.Since(start).String()).
			Str(userAgent, c.Get(fiber.HeaderUserAgent)).
			Logger()

		switch {
		case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
			reqLogger.Warn().Msg(httpReq)
		case code >= http.StatusInternalServerError:
			reqLogger.Error().Msg(httpReq)
		default:
			reqLogger.Info().Msg(httpReq)
		}

		return nil
	}
}

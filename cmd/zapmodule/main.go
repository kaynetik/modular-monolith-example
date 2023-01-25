package main

import (
	"fmt"
	"time"

	"github.com/kaynetik/modular-monolith-example/internal/app/api/handlers"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/config"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/server"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/storage"
	postgresdb "github.com/kaynetik/modular-monolith-example/internal/pkg/storage/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	time.Local = time.UTC
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	conf, err := config.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to set up the Config: %w", err))
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if conf.LogLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	apiDB, err := postgresdb.New(conf.DatabaseURL)
	if err != nil {
		log.Error().Err(err).Msg("")

		return
	}

	var (
		repo   = storage.New(apiDB, &conf)
		engine server.Engine
	)

	engine.New(&conf.Server)

	err = handlers.
		Attach(engine, repo).
		Serve()
	if err != nil {
		log.Fatal().Err(err).Msg("server was not started")
	}

	log.Error().Err(err).Msg("server run into an issue, and was stopped")
}

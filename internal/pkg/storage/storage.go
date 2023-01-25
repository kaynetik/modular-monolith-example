package storage

import (
	"github.com/uptrace/bun"

	"github.com/kaynetik/modular-monolith-example/internal/pkg/config"
)

type Repository struct {
	db bun.IDB
	// redis  *redisdb.Conn

	conf *config.Config
}

func New(dbConn bun.IDB, conf *config.Config) *Repository {
	return &Repository{dbConn, conf}
}

func (r *Repository) GetConfig() *config.Config {
	return r.conf
}

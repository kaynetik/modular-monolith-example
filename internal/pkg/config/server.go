package config

import (
	"strings"
	"time"

	"github.com/kaynetik/modular-monolith-example/internal/pkg/env"
)

type Stage string

const (
	StageProduction  Stage = "PRODUCTION"
	StageStaging     Stage = "STAGING"
	StageIntegration Stage = "INT"
	StageTesting     Stage = "DEVELOP"
)

func (s Stage) Is(stage string) bool {
	return strings.EqualFold(string(s), stage)
}

func CurrentStage() Stage {
	stage := env.Get(env.Stage)
	stage = strings.TrimSpace(stage)
	stage = strings.ToUpper(stage)

	switch Stage(stage) {
	case StageProduction:
		return StageProduction
	case StageStaging:
		return StageStaging
	case StageIntegration:
		return StageIntegration
	case StageTesting:
		return StageTesting
	}

	return ""
}

const (
	defaultServerReadWriteTimeout = time.Minute
)

type Server struct {
	RunMode      Stage
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	AppName      string
}

func readServerConfig() Server {
	return Server{
		RunMode:      CurrentStage(),
		HTTPPort:     env.Get("API_PORT"),
		ReadTimeout:  defaultServerReadWriteTimeout,
		WriteTimeout: defaultServerReadWriteTimeout,
		AppName:      env.GetOr(env.AppName, "local"),
	}
}

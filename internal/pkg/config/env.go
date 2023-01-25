package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/env"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Server      Server
	LogLevel    string // LogLevel - add an enum for a few levels, to ensure stability.
	DatabaseURL string
	RedisURL    string
	APIPort     string
	APIVer      string
	APIBase     string
	AdminEmail  string
	HMACSecKey  []byte
	appSecret   string
	adminPass   string

	PrivKeyPEM []byte
	PubKeyPEM  []byte
}

type EmailConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
}

func ReadConfig() (config Config, err error) {
	PreloadEnvironment()

	privateECKeyPath := env.Get("PRIVATE_EC_KEY")
	publicECKeyPath := env.Get("PUBLIC_EC_KEY")

	privateKey, err := os.ReadFile(privateECKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed reading private key from pem file")
	}

	publicKey, err := os.ReadFile(publicECKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed reading public key from pem file")
	}

	httpPort := env.GetOr(env.APIPort, "9450")

	config.RedisURL = env.GetOr(env.RedisURL, "host=redis://localhost:6379")
	config.LogLevel = env.GetOr(env.LogLevel, "info")
	config.Server = readServerConfig()
	config.APIPort = env.GetOr(env.APIPort, httpPort)
	config.APIVer = env.GetOr(env.APIVer, "/api/v1")
	config.APIBase = env.GetOr(env.APIBase, "http://localhost:"+config.APIPort)

	config.appSecret = env.GetOr(env.AppSecret, "huge-sec")
	config.HMACSecKey = []byte(config.appSecret)

	config.AdminEmail = env.GetOr(env.AdminEmail, "contact@decantera.dev")
	config.adminPass = env.GetOr(env.AdminPass, "dsafkl9i3jhnxS@$#")

	config.PrivKeyPEM = privateKey
	config.PubKeyPEM = publicKey

	log.Info().Msg("Read config successfully")

	return config, err
}

func PreloadEnvironment() {
	if err := godotenv.Load(); err != nil {
		// Don't return error as the config might be not in file, but in env vars already
		log.Error().Err(err).Msg("Failed to load env vars from env file.")
	}
}

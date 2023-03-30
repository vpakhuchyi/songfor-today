package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/exp/slog"
)

type Config struct {
	Authorization     Authorization
	LogLevel          string `envconfig:"LOG_LEVEL"`
	ProjectID         string `envconfig:"PROJECT_ID"`
	DeezerAppSecret   string `envconfig:"DEEZER_APP_SECRET"`
	DeezerAccessToken string `envconfig:"DEEZER_ACCESS_TOKEN"`
}

type Authorization struct {
	AppID       string `envconfig:"APP_ID"`
	Permissions string `envconfig:"PERMISSIONS"`
	RedirectURI string `envconfig:"REDIRECT_URI"`
	LoginURL    string `envconfig:"LOGIN_URL"`
	TokenURL    string `envconfig:"TOKEN_URL"`
}

// MustReadConfiguration reads configuration from env variables.
func MustReadConfiguration() Config {
	var cfg Config
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Failed to load ENV variables from .env file", "err", err)
		panic(err)
	}

	envconfig.MustProcess("", &cfg)

	slog.Info("Successfully loaded the configuration")

	return cfg
}

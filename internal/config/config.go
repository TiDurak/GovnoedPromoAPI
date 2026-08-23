package config

import (
	"errors"
	"os"
)

type Config struct {
	HTTPPort     string
	DatabasePath string

	DiscordClientID     string
	DiscordClientSecret string
	DiscordRedirectURL  string
	PromoReward         int
}

func Load() Config {
	return Config{
		HTTPPort: "8080",

		DatabasePath: getEnv("PROMO_DATABASE_PATH", "./data/database.db"),

		DiscordClientID:     getEnv("DISCORD_CLIENT_ID", ""),
		DiscordClientSecret: getEnv("DISCORD_CLIENT_SECRET", ""),
		DiscordRedirectURL:  getEnv("DISCORD_REDIRECT_URL", ""),
		PromoReward:         750,
	}
}

func (c Config) Validate() error {
	if c.HTTPPort == "" {
		return errors.New("HTTP port is not configured")
	}

	if c.DatabasePath == "" {
		return errors.New("database path is not configured")
	}

	if c.DiscordClientID == "" {
		return errors.New("Discord client ID is not configured")
	}

	if c.DiscordClientSecret == "" {
		return errors.New("Discord client secret is not configured")
	}

	if c.DiscordRedirectURL == "" {
		return errors.New("Discord redirect URL is not configured")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

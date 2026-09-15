// Package config handles loading and validating application configuration
// from environment variables.
package config

import (
	"errors"
	"os"
)

// ErrMissingToken is returned when the TMDB API token is not set.
var ErrMissingToken = errors.New(
	"TMDB_API_TOKEN is not set.\n\n" +
		"Get a free API Read Access Token at:\n" +
		"  https://www.themoviedb.org/settings/api\n\n" +
		"Then export it:\n" +
		"  export TMDB_API_TOKEN=\"your_token_here\"",
)

// Config holds the application configuration.
type Config struct {
	APIToken string
}

// Load reads configuration from environment variables and validates it.
func Load() (Config, error) {
	token := os.Getenv("TMDB_API_TOKEN")
	if token == "" {
		return Config{}, ErrMissingToken
	}

	return Config{APIToken: token}, nil
}

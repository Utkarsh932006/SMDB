package config_test

import (
	"errors"
	"testing"

	"github.com/Utkarsh932006/SMDB/internal/config"
)

func TestLoad_Success(t *testing.T) {
	t.Setenv("TMDB_API_TOKEN", "test_token_123")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.APIToken != "test_token_123" {
		t.Errorf("expected token 'test_token_123', got '%s'", cfg.APIToken)
	}
}

func TestLoad_MissingToken(t *testing.T) {
	t.Setenv("TMDB_API_TOKEN", "")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for missing token, got nil")
	}

	if !errors.Is(err, config.ErrMissingToken) {
		t.Errorf("expected ErrMissingToken, got %v", err)
	}
}

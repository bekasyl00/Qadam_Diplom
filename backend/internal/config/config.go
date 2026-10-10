package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port            string
	DatabaseURL     string
	AIMode          string
	AIProvider      string
	AIBaseURL       string
	OpenAIKey       string
	OpenAIModel     string
	TranscribeModel string
	DemoMode        bool
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("could not parse .env")
	}
	provider := valueOrDefault("AI_PROVIDER", "openai")
	key := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	model := valueOrDefault("OPENAI_MODEL", "gpt-4.1-mini")
	transcribeModel := valueOrDefault("OPENAI_TRANSCRIBE_MODEL", "gpt-4o-mini-transcribe")
	baseURL := "https://api.openai.com/v1"
	if provider == "groq" {
		key = strings.TrimSpace(os.Getenv("GROQ_API_KEY"))
		model = valueOrDefault("AI_MODEL", "openai/gpt-oss-20b")
		transcribeModel = valueOrDefault("AI_TRANSCRIBE_MODEL", "whisper-large-v3-turbo")
		baseURL = valueOrDefault("AI_BASE_URL", "https://api.groq.com/openai/v1")
	}
	cfg := Config{
		Port:            valueOrDefault("PORT", "8080"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		AIMode:          valueOrDefault("AI_MODE", "fallback"),
		AIProvider:      provider,
		AIBaseURL:       strings.TrimRight(baseURL, "/"),
		OpenAIKey:       key,
		OpenAIModel:     model,
		TranscribeModel: transcribeModel,
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("PORT must be 1-65535")
	}
	cfg.DemoMode, err = strconv.ParseBool(valueOrDefault("DEMO_MODE", "true"))
	if err != nil {
		return Config{}, fmt.Errorf("DEMO_MODE must be true or false")
	}
	if cfg.AIMode != "live" && cfg.AIMode != "fallback" {
		return Config{}, fmt.Errorf("AI_MODE must be live or fallback")
	}
	if cfg.AIProvider != "openai" && cfg.AIProvider != "groq" {
		return Config{}, fmt.Errorf("AI_PROVIDER must be openai or groq")
	}
	if cfg.AIMode == "live" && cfg.OpenAIKey == "" {
		return Config{}, fmt.Errorf("AI provider key is required for AI_MODE=live")
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func valueOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

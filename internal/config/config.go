package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Config defines a configuration of secrets and environment variables
// necessary for the app to function.
type Config struct {
	Port          string
	BaseURL       string
	EBirdAPIToken string
	DB            PostgresConfig
}

// PostgresConfig defines a configuration of credentials for connecting to
// a PostgreSQL database.
type PostgresConfig struct {
	Password string
	User     string
	DB       string
	Port     string
}

// LoadConfig returns a pointer to a Config or an error.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:          os.Getenv("PORT"),
		BaseURL:       os.Getenv("API_BASE_URL"),
		EBirdAPIToken: os.Getenv("EBIRD_API_KEY"),
		DB: PostgresConfig{
			Password: os.Getenv("POSTGRES_PASSWORD"),
			User:     os.Getenv("POSTGRES_USER"),
			DB:       os.Getenv("POSTGRES_DB"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}
	return cfg, nil
}

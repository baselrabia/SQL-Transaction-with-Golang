package config

import (
	"os"
	"strconv"
)

// Config represents the PostgreSQL connection configuration
type EnvConfig struct {
	PgHost     string
	PgPort     int
	PgUser     string
	PgPassword string
	PgDatabase string
	PgSslMode  string
}

// NewConfig initializes a new Config instance with values from environment variables
func NewEnvConfig() EnvConfig {
	return EnvConfig{
		PgHost:     getEnv("PG_HOST", "localhost"),
		PgPort:     getEnvAsInt("PG_PORT", 5432),
		PgUser:     getEnv("PG_USER", ""),
		PgPassword: getEnv("PG_PASSWORD", ""),
		PgDatabase: getEnv("PG_DATABASE", ""),
		PgSslMode:  getEnv("PG_SSL_MODE", "disable"),
	}
}

// getEnv returns the value of the specified environment variable,
// or the default value if the environment variable is not set
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt returns the value of the specified environment variable as an integer,
// or the default value if the environment variable is not set or cannot be parsed as an integer
func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if valueInt, err := strconv.Atoi(valueStr); err == nil {
			return valueInt
		}
	}
	return defaultValue
}

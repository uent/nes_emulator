package config

// AppConfig holds the application configuration
type AppConfig struct {
	Debug   bool
	Version string
}

// NewConfig returns a new configuration with default values
func NewConfig() *AppConfig {
	return &AppConfig{
		Debug:   false,
		Version: "0.1.0",
	}
}
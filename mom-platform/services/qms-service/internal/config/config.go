package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the QMS service.
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	NATS     NATSConfig     `yaml:"nats"`
	Redis    RedisConfig    `yaml:"redis"`
}

// ServerConfig holds gRPC server configuration.
type ServerConfig struct {
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}

// DatabaseConfig holds PostgreSQL connection parameters.
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// DSN returns the PostgreSQL connection string.
func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// NATSConfig holds NATS connection parameters.
type NATSConfig struct {
	URL string `yaml:"url"`
}

// RedisConfig holds Redis connection parameters.
type RedisConfig struct {
	Addr string `yaml:"addr"`
}

// Load reads configuration from a YAML file at the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	if cfg.Server.Port == 0 {
		cfg.Server.Port = 50053
	}
	if cfg.Server.Name == "" {
		cfg.Server.Name = "qms-service"
	}
	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}

	applyEnvOverrides(&cfg)

	return &cfg, nil
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("DB_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &cfg.Database.Port)
	}
	if v := os.Getenv("DB_USER"); v != "" {
		cfg.Database.User = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		cfg.Database.DBName = v
	}
	if v := os.Getenv("NATS_URL"); v != "" {
		cfg.NATS.URL = v
	}
	if v := os.Getenv("REDIS_ADDR"); v != "" {
		cfg.Redis.Addr = v
	}
}

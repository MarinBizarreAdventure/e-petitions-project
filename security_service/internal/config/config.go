package config

import (
	"os"
	"path/filepath"

	"github.com/gookit/slog"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GrpcPort int      `yaml:"grpc_port"`
	Database Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
	Token    Token    `yaml:"token"`
	Rabbit   Rabbit   `yaml:"rabbit"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	RedisDB  int    `yaml:"db"`
}

type Token struct {
	TKey  string `yaml:"T_KEY"`
	RTKey string `yaml:"RT_KEY"`
}

type Rabbit struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() *Config {
	var cfg *Config
	wd, err := os.Getwd()
	if err != nil {
		slog.Fatalf("Failed to get working directory: %v", err)
	}
	configPath := filepath.Join(wd, "config.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return cfg
}

package config

import (
	"github.com/BurntSushi/toml"
	db "github.com/evtepo/auth/internal/infrastructure/db/config"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type AppConfig struct {
	Mode string `toml:"mode"`
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Config struct {
	App AppConfig   `toml:"app"`
	Db  db.DBConfig `toml:"db"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(path string, config *Config) error {
	absolutePath, wdErr := os.Getwd()
	if wdErr != nil {
		return wdErr
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		return errEnv
	}

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = path
	}

	filePath := filepath.Join(absolutePath, cfgPath)

	_, err := toml.DecodeFile(filePath, config)
	return err
}

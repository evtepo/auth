package config

import (
	"fmt"
)

type DBConfig struct {
	User      string `toml:"user"`
	Password  string `toml:"password"`
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	Database  string `toml:"database"`
	Migration bool   `toml:"migration"`
}

func (conf *DBConfig) FullDNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port,
	)
}

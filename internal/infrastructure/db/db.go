package db

import (
	"github.com/evtepo/auth/internal/infrastructure/db/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config *config.DBConfig) *gorm.DB {
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(config.FullDNS()), gormConfig)
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection(db *gorm.DB) error {
	connection, err := db.DB()
	if err != nil {
		return err
	}

	connection.Close()
	return nil
}

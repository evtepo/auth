package db

import (
	"github.com/evtepo/auth/internal/infrastructure/db/config"
	migration "github.com/evtepo/auth/internal/infrastructure/db/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config *config.DBConfig) *gorm.DB {
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(config.FullDNS()), gormConfig)
	if err == nil {
		if config.Migration {
			migration.Migrate(db)
		}
		return db
	}

	panic(err.Error())
}

func CloseConnection(db *gorm.DB) error {
	connection, err := db.DB()
	if err != nil {
		return err
	}

	connection.Close()
	return nil
}

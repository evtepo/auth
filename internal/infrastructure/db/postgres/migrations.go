package postgres

import (
	"gorm.io/gorm"
	"github.com/evtepo/auth/internal/infrastructure/db/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Role{})
}

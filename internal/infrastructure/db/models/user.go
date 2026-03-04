package models

type User struct {
	Base
	Username string
	Email    string
	Password string
	RoleID   int
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsAdmin  bool `gorm:"default:false"`
}

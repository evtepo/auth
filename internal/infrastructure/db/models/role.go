package models

type Role struct {
	Base
	Name  string
	Users []User `gorm:"foreignKey:RoleID"`
}

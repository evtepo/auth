package user

import (
	"github.com/evtepo/auth/internal/domain/base"
)

type User struct {
	base.BaseModel
	Username  Username
	Email     Email
	Password  Password
	RoleID    uint
	IsAdmin   bool
}

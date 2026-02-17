package user

import (
	"github.com/evtepo/auth/internal/domain/role"
)

type User struct {
	ID       int
	Usename  Usename
	Email    Email
	Password Password
	RoleID   []role.Role
	IsDel    bool
	IsAdmin  bool
}

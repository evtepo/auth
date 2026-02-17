package user

import (
	"context"

	. "github.com/evtepo/auth/internal/domain/role"
)

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	ChangeRole(ctx context.Context, user *User, role *Role) error
	FindById(ctx context.Context, id int) (*User, error)
	FindByUsername(ctx context.Context, username Usename) (*User, error)
	GetUserInfo(ctx context.Context, id int) (*User, error)
}
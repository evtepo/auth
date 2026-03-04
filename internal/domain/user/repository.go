package user

import (
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	ChangeRole(ctx context.Context, user *User, roleId uint) error
	FindById(ctx context.Context, id uint) (*User, error)
	FindByUsername(ctx context.Context, username Username) (*User, error)
	GetUserInfo(ctx context.Context, req *User) (*User, error)
	DeleteUser(ctx context.Context, userId uint) error
}
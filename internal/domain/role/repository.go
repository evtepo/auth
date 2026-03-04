package role

import (
	"context"
)

type RoleRepository interface {
	Create(ctx context.Context, role *Role) error
	FindById(ctx context.Context, roleId uint) (*Role, error)
	FindByName(ctx context.Context, name RoleName) (*Role, error)
	Delete(ctx context.Context, roleId uint) error
}
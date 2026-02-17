package postgres

import (
	"context"

	"github.com/evtepo/auth/internal/domain/user"
)

type UserRepo struct{
	user.UserRepository
}

func NewRepository() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Register(ctx context.Context, user *user.User) error {
	return nil
}

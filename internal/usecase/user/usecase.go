package user

import (
	"context"

	userDomain "github.com/evtepo/auth/internal/domain/user"
)

type Usecase struct {
	repo userDomain.UserRepository
}

func NewUserUsecase(repository userDomain.UserRepository) *Usecase {
	return &Usecase{repo: repository}
}

func (u *Usecase) Register(ctx context.Context, user *userDomain.User) error {
	return nil
}

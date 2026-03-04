package user

import (
	"context"
	"errors"

	roleDomain "github.com/evtepo/auth/internal/domain/role"
	userDomain "github.com/evtepo/auth/internal/domain/user"
)

type Usecase struct {
	userRepo userDomain.UserRepository
	roleRepo roleDomain.RoleRepository
}

func NewUserUsecase(userRepo userDomain.UserRepository, roleRepo roleDomain.RoleRepository) *Usecase {
	return &Usecase{userRepo: userRepo, roleRepo: roleRepo}
}

func (u *Usecase) Register(
	ctx context.Context,
	email userDomain.Email,
	username userDomain.Username,
	password userDomain.Password,
) error {
	roleName := roleDomain.RoleName("Member")
	roleId, err := u.roleRepo.FindByName(ctx, roleName)
	if err != nil {
		return errors.New("Role not found")
	}
	hashedPassword := password
	entity := &userDomain.User{
		Email: email,
		Username: username,
		Password: hashedPassword,
		RoleID: roleId.ID,
	}
	registerErr := u.userRepo.Register(ctx, entity)
	if registerErr != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

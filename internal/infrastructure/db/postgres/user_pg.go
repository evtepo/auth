package postgres

import (
	"context"

	"github.com/evtepo/auth/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepo struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Register(ctx context.Context, user *user.User) error {
	user.IsAdmin = false
	newUser := r.db.Create(&user)
	return newUser.Error
}

func (r *UserRepo) ChangeRole(ctx context.Context, user *user.User, roleId uint) error {
	user.RoleID = roleId
	result := r.db.Save(&user)
	return result.Error
}

func (r * UserRepo) FindById(ctx context.Context, id uint) (*user.User, error) {
	var foundUser user.User
	result := r.db.First(&foundUser, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}

func (r *UserRepo) FindByUsername(ctx context.Context, username user.Username) (*user.User, error) {
	var foundUser user.User
	result := r.db.Where("username = ?", username).First(&foundUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}

func (r *UserRepo) GetUserInfo(ctx context.Context, req *user.User) (*user.User, error) {
	var foundUser user.User
	result := r.db.Where("id = ?", req.ID).First(&foundUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, userId uint) error {
	result := r.db.Delete(&user.User{}, userId)
	return result.Error
}
package postgres

import (
	"context"

	"github.com/evtepo/auth/internal/domain/role"
	"gorm.io/gorm"
)

type RoleRepo struct{
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) Create(ctx context.Context, role *role.Role) error {
	newRole := r.db.Create(&role)
	return newRole.Error
}

func (r *RoleRepo) FindByName(ctx context.Context, name role.RoleName) (*role.Role, error) {
	var foundRole role.Role
	result := r.db.Where("name = ?", name).First(&foundRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundRole, nil
}

func (r *RoleRepo) FindById(ctx context.Context, roleId uint) (*role.Role, error) {
	var foundRole role.Role
	result := r.db.First(&foundRole, roleId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundRole, nil
}

func (r *RoleRepo) Delete(ctx context.Context, roleId uint) error {
	result := r.db.Delete(&role.Role{}, roleId)
	return result.Error
}

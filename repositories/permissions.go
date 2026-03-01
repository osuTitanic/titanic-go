package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type UserPermissionRepository struct {
	db *gorm.DB
}

func NewUserPermissionRepository(db *gorm.DB) *UserPermissionRepository {
	return &UserPermissionRepository{db: db}
}

func (r *UserPermissionRepository) Create(permission *schemas.UserPermission) error {
	return r.db.Create(permission).Error
}

func (r *UserPermissionRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.UserPermission{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserPermissionRepository) Delete(permission *schemas.UserPermission) error {
	return r.db.Delete(permission).Error
}

func (r *UserPermissionRepository) ById(id int) (*schemas.UserPermission, error) {
	var permission schemas.UserPermission
	err := r.db.Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *UserPermissionRepository) ManyByUserId(userId int) ([]*schemas.UserPermission, error) {
	var permissions []*schemas.UserPermission
	err := r.db.Where("user_id = ?", userId).Find(&permissions).Error
	return permissions, err
}

type GroupPermissionRepository struct {
	db *gorm.DB
}

func NewGroupPermissionRepository(db *gorm.DB) *GroupPermissionRepository {
	return &GroupPermissionRepository{db: db}
}

func (r *GroupPermissionRepository) Create(permission *schemas.GroupPermission) error {
	return r.db.Create(permission).Error
}

func (r *GroupPermissionRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.GroupPermission{}).Where("id = ?", id).Updates(updates).Error
}

func (r *GroupPermissionRepository) Delete(permission *schemas.GroupPermission) error {
	return r.db.Delete(permission).Error
}

func (r *GroupPermissionRepository) ById(id int) (*schemas.GroupPermission, error) {
	var permission schemas.GroupPermission
	err := r.db.Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *GroupPermissionRepository) ManyByGroupId(groupId int) ([]*schemas.GroupPermission, error) {
	var permissions []*schemas.GroupPermission
	err := r.db.Where("group_id = ?", groupId).Find(&permissions).Error
	return permissions, err
}

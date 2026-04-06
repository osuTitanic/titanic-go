package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
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

func (r *UserPermissionRepository) Delete(permission *schemas.UserPermission) error {
	return r.db.Delete(permission).Error
}

func (r *UserPermissionRepository) Update(updates *schemas.UserPermission, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}

func (r *UserPermissionRepository) ById(id int, preload ...string) (*schemas.UserPermission, error) {
	var permission schemas.UserPermission
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *UserPermissionRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.UserPermission, error) {
	var permissions []*schemas.UserPermission
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&permissions).Error
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

func (r *GroupPermissionRepository) Delete(permission *schemas.GroupPermission) error {
	return r.db.Delete(permission).Error
}

func (r *GroupPermissionRepository) Update(updates *schemas.GroupPermission, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *GroupPermissionRepository) ById(id int, preload ...string) (*schemas.GroupPermission, error) {
	var permission schemas.GroupPermission
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *GroupPermissionRepository) ManyByGroupId(groupId int, preload ...string) ([]*schemas.GroupPermission, error) {
	var permissions []*schemas.GroupPermission
	err := Preloaded(r.db, preload).Where("group_id = ?", groupId).Find(&permissions).Error
	return permissions, err
}

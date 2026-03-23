package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(group *schemas.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepository) Delete(group *schemas.Group) error {
	return r.db.Delete(group).Error
}

func (r *GroupRepository) Update(updates *schemas.Group, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *GroupRepository) ById(id int, preload ...string) (*schemas.Group, error) {
	var group schemas.Group
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepository) Many(includeHidden bool, preload ...string) ([]*schemas.Group, error) {
	query := Preloaded(r.db, preload).Model(&schemas.Group{})
	if !includeHidden {
		query = query.Where("hidden = ?", false)
	}

	var groups []*schemas.Group
	err := query.Order("id ASC").Find(&groups).Error
	return groups, err
}

type GroupEntryRepository struct {
	db *gorm.DB
}

func NewGroupEntryRepository(db *gorm.DB) *GroupEntryRepository {
	return &GroupEntryRepository{db: db}
}

func (r *GroupEntryRepository) Create(entry *schemas.GroupEntry) error {
	return r.db.Create(entry).Error
}

func (r *GroupEntryRepository) Update(updates *schemas.GroupEntry, columns ...string) (int64, error) {
	var result *gorm.DB
	if len(columns) == 0 {
		result = r.db.Save(&updates)
	} else {
		result = r.db.Model(&updates).Select(columns).Updates(&updates)
	}
	return result.RowsAffected, result.Error
}

func (r *GroupEntryRepository) Delete(entry *schemas.GroupEntry) error {
	return r.db.Delete(entry).Error
}

func (r *GroupEntryRepository) ByUserAndGroup(userId int, groupId int, preload ...string) (*schemas.GroupEntry, error) {
	var entry schemas.GroupEntry
	err := Preloaded(r.db, preload).Where("user_id = ? AND group_id = ?", userId, groupId).First(&entry).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *GroupEntryRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.GroupEntry, error) {
	var entries []*schemas.GroupEntry
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&entries).Error
	return entries, err
}

func (r *GroupEntryRepository) ManyByGroupId(groupId int, preload ...string) ([]*schemas.GroupEntry, error) {
	var entries []*schemas.GroupEntry
	err := Preloaded(r.db, preload).Where("group_id = ?", groupId).Find(&entries).Error
	return entries, err
}

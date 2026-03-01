package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
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

func (r *GroupRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.Group{}).Where("id = ?", id).Updates(updates).Error
}

func (r *GroupRepository) Delete(group *schemas.Group) error {
	return r.db.Delete(group).Error
}

func (r *GroupRepository) ById(id int) (*schemas.Group, error) {
	var group schemas.Group
	err := r.db.Where("id = ?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepository) Many(includeHidden bool) ([]*schemas.Group, error) {
	query := r.db.Model(&schemas.Group{})
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

func (r *GroupEntryRepository) Update(userId int, groupId int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.GroupEntry{}).Where("user_id = ? AND group_id = ?", userId, groupId).Updates(updates).Error
}

func (r *GroupEntryRepository) Delete(entry *schemas.GroupEntry) error {
	return r.db.Delete(entry).Error
}

func (r *GroupEntryRepository) ByUserAndGroup(userId int, groupId int) (*schemas.GroupEntry, error) {
	var entry schemas.GroupEntry
	err := r.db.Where("user_id = ? AND group_id = ?", userId, groupId).First(&entry).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *GroupEntryRepository) ManyByUserId(userId int) ([]*schemas.GroupEntry, error) {
	var entries []*schemas.GroupEntry
	err := r.db.Where("user_id = ?", userId).Find(&entries).Error
	return entries, err
}

func (r *GroupEntryRepository) ManyByGroupId(groupId int) ([]*schemas.GroupEntry, error) {
	var entries []*schemas.GroupEntry
	err := r.db.Where("group_id = ?", groupId).Find(&entries).Error
	return entries, err
}

package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type RelationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) *RelationshipRepository {
	return &RelationshipRepository{db: db}
}

func (r *RelationshipRepository) Create(relationship *schemas.Relationship) error {
	return r.db.Create(relationship).Error
}

func (r *RelationshipRepository) Update(userId int, targetId int, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&schemas.Relationship{}).Where("user_id = ? AND target_id = ?", userId, targetId).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *RelationshipRepository) Delete(relationship *schemas.Relationship) error {
	return r.db.Delete(relationship).Error
}

func (r *RelationshipRepository) ByUserAndTarget(userId int, targetId int, preload ...string) (*schemas.Relationship, error) {
	var relationship schemas.Relationship
	err := Preloaded(r.db, preload).Where("user_id = ? AND target_id = ?", userId, targetId).First(&relationship).Error
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

func (r *RelationshipRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Relationship, error) {
	var relationships []*schemas.Relationship
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&relationships).Error
	return relationships, err
}

func (r *RelationshipRepository) ManyByTargetId(targetId int, preload ...string) ([]*schemas.Relationship, error) {
	var relationships []*schemas.Relationship
	err := Preloaded(r.db, preload).Where("target_id = ?", targetId).Find(&relationships).Error
	return relationships, err
}

func (r *RelationshipRepository) CountByUserId(userId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.Relationship{}).Where("user_id = ?", userId).Count(&count).Error
	return int(count), err
}

func (r *RelationshipRepository) CountByTargetId(targetId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.Relationship{}).Where("target_id = ?", targetId).Count(&count).Error
	return int(count), err
}

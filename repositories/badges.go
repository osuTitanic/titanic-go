package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type BadgeRepository struct {
	db *gorm.DB
}

func NewBadgeRepository(db *gorm.DB) *BadgeRepository {
	return &BadgeRepository{db: db}
}

func (r *BadgeRepository) Create(badge *schemas.Badge) error {
	return r.db.Create(badge).Error
}

func (r *BadgeRepository) Update(id int, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&schemas.Badge{}).Where("id = ?", id).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *BadgeRepository) Delete(badge *schemas.Badge) error {
	return r.db.Delete(badge).Error
}

func (r *BadgeRepository) ById(id int, preload ...string) (*schemas.Badge, error) {
	var badge schemas.Badge
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&badge).Error
	if err != nil {
		return nil, err
	}
	return &badge, nil
}

func (r *BadgeRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Badge, error) {
	var badges []*schemas.Badge
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Order("created DESC").Find(&badges).Error
	return badges, err
}

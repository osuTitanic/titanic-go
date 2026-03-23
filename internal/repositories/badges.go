package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
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

func (r *BadgeRepository) Delete(badge *schemas.Badge) error {
	return r.db.Delete(badge).Error
}

func (r *BadgeRepository) Update(updates *schemas.Badge, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
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

package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type AchievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository(db *gorm.DB) *AchievementRepository {
	return &AchievementRepository{db: db}
}

func (r *AchievementRepository) Create(achievement *schemas.Achievement) error {
	return r.db.Create(achievement).Error
}

func (r *AchievementRepository) Delete(achievement *schemas.Achievement) error {
	return r.db.Delete(achievement).Error
}

func (r *AchievementRepository) Update(updates *schemas.Achievement, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *AchievementRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Achievement, error) {
	var achievements []*schemas.Achievement
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&achievements).Error
	return achievements, err
}

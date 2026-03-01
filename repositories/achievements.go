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

func (r *AchievementRepository) Update(userId int, name string, updates map[string]interface{}) error {
	return r.db.Model(&schemas.Achievement{}).Where("user_id = ? AND name = ?", userId, name).Updates(updates).Error
}

func (r *AchievementRepository) Delete(achievement *schemas.Achievement) error {
	return r.db.Delete(achievement).Error
}

func (r *AchievementRepository) ManyByUserId(userId int) ([]*schemas.Achievement, error) {
	var achievements []*schemas.Achievement
	err := r.db.Where("user_id = ?", userId).Find(&achievements).Error
	return achievements, err
}

package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type BeatmapModdingRepository struct {
	db *gorm.DB
}

func NewBeatmapModdingRepository(db *gorm.DB) *BeatmapModdingRepository {
	return &BeatmapModdingRepository{db: db}
}

func (r *BeatmapModdingRepository) Create(modding *schemas.BeatmapModding) error {
	return r.db.Create(modding).Error
}

func (r *BeatmapModdingRepository) Delete(modding *schemas.BeatmapModding) error {
	return r.db.Delete(modding).Error
}

func (r *BeatmapModdingRepository) Update(updates *schemas.BeatmapModding, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}

func (r *BeatmapModdingRepository) TotalKudosuByUser(userId int) (int, error) {
	var total int
	err := r.db.Model(&schemas.BeatmapModding{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("target_id = ?", userId).
		Scan(&total).Error
	return total, err
}
